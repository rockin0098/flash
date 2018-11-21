package tlservice

import (
	"github.com/rockin0098/meow/proto/mtproto"
	"github.com/rockin0098/meow/server/model"
	"github.com/rockin0098/meow/server/service"
)

// TL_messages_getPinnedDialogs
func (s *TLService) TL_messages_getPinnedDialogs_Process(csess *service.ClientSession, object mtproto.TLObject) (mtproto.TLObject, error) {
	Log.Infof("entering... client sessid = %v", csess.ClientSessionID)

	// tlobj := object
	// tl := tlobj.(*mtproto.TL_messages_getPinnedDialogs)

	userid := csess.GetUserID()

	dialogs := s.Dao.UserDialogDao.GetPinnedDialogs(userid)
	tldialogs := model.DialogList_to_TL_dialogList(dialogs)
	peerDialogs := mtproto.New_TL_messages_peerDialogs()

	messageIdList := []int32{}
	userIdList := []int64{userid}
	chatIdList := []int32{}

	for _, tld := range tldialogs {
		tldialog := tld.(*mtproto.TL_dialog)
		messageIdList = append(messageIdList, tldialog.Get_top_message())
		peer := tldialog.Get_peer()
		if peer.ClassID() == mtproto.TL_CLASS_peerUser {
			userIdList = append(userIdList, int64(peer.(*mtproto.TL_peerUser).Get_user_id()))
		} else if peer.ClassID() == mtproto.TL_CLASS_peerChat {
			chatIdList = append(chatIdList, peer.(*mtproto.TL_peerChat).Get_chat_id())
		} else if peer.ClassID() == mtproto.TL_CLASS_peerChannel {
			// do nothing
		} else {
			panic("invalid peer")
		}

		peerDialogs.M_dialogs = append(peerDialogs.M_dialogs, tldialog)
	}

	if len(messageIdList) > 0 {
		messages := s.Dao.MessageDao.GetMessagesByIDList(userid, messageIdList)
		peerDialogs.M_messages = model.MessageList_to_TL_messageList(messages)
	}

	users := s.Dao.UserDao.GetUsersBySelfAndIDList(userid, userIdList)
	peerDialogs.M_users = users

	if len(chatIdList) > 0 {
		peerDialogs.M_chats = s.Dao.ChatDao.GetChatListBySelfAndIDList(userid, chatIdList)
	}

	// state 机制 未处理
	peerDialogs.M_state = mtproto.New_TL_updates_state()

	return peerDialogs, nil
}
