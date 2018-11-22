package tlservice

import (
	"fmt"
	"math"

	"github.com/rockin0098/meow/server/model"

	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/service"
)

type DialogItems struct {
	MessageIDList       []int32
	ChannelMessageIDMap map[int32]int32
	UserIDList          []int32
	ChatIDList          []int32
	ChannelIDList       []int32
}

// TL_messages_getDialogs
func (s *TLService) TL_messages_getDialogs_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	tlobj := object
	tl := tlobj.(*mtproto.TL_messages_getDialogs)

	offsetID := tl.Get_offset_id()
	if offsetID == 0 {
		offsetID = math.MaxInt32
	}

	userid := csess.GetUserID()
	limit := tl.Get_limit()

	dialogs := s.Dao.UserDialogDao.GetDialogsByOffsetID(userid, false, offsetID, limit)
	tldialogs := model.DialogList_to_TL_dialogList(dialogs)
	items := PickAllIDListByDialogs(tldialogs)
	mss := s.Dao.MessageDao.GetMessagesByIDList(userid, items.MessageIDList)
	tlmss := model.MessageList_to_TL_messageList(mss)
	for k, v := range items.ChannelMessageIDMap {
		m := s.Dao.MessageDao.GetChannelMessage(k, v)
		if m != nil {
			tlmss = append(tlmss, model.Message_to_TL_message(m))
		}
	}

	users := s.Dao.UserDao.GetUsersBySelfAndIDList(userid, items.UserIDList)
	chats := s.Dao.ChatDao.GetChatListBySelfAndIDList(userid, items.ChatIDList)
	// todo : get channel chat

	messageDialogs := &mtproto.TL_messages_dialogs{
		M_dialogs:  tldialogs,
		M_messages: tlmss,
		M_users:    users,
		M_chats:    chats,
	}

	return messageDialogs, nil
}

func PickAllIDListByDialogs(tldialogs []mtproto.TLObject) (items *DialogItems) {
	items = &DialogItems{}

	for _, td := range tldialogs {

		d := td.(*mtproto.TL_dialog)

		peer := d.Get_peer()

		switch peer.(type) {
		case *mtproto.TL_peerUser:
			items.MessageIDList = append(items.MessageIDList, d.Get_top_message())
			items.UserIDList = append(items.UserIDList, (peer.(*mtproto.TL_peerUser).Get_user_id()))
		case *mtproto.TL_peerChat:
			items.MessageIDList = append(items.MessageIDList, d.Get_top_message())
			items.UserIDList = append(items.UserIDList, (peer.(*mtproto.TL_peerChat).Get_chat_id()))
		case *mtproto.TL_peerChannel:
			p := peer.(*mtproto.TL_peerChannel)
			items.ChannelMessageIDMap[p.Get_channel_id()] = d.Get_top_message()
			items.UserIDList = append(items.UserIDList, (p.Get_channel_id()))
		default:
			panic(fmt.Sprintf("invalid type = %T", d))
		}
	}

	return
}
