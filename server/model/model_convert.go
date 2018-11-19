package model

import (
	"fmt"
	"time"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
)

const (
	MESSAGE_TYPE_UNKNOWN         = 0
	MESSAGE_TYPE_MESSAGE_EMPTY   = 1
	MESSAGE_TYPE_MESSAGE         = 2
	MESSAGE_TYPE_MESSAGE_SERVICE = 3
)
const (
	MESSAGE_BOX_TYPE_INCOMING = 0
	MESSAGE_BOX_TYPE_OUTGOING = 1
)

const (
	PTS_UNKNOWN             = 0
	PTS_MESSAGE_OUTBOX      = 1
	PTS_MESSAGE_INBOX       = 2
	PTS_READ_HISTORY_OUTBOX = 3
	PTS_READ_HISTORY_INBOX  = 4
)

func User_to_TL_user(user *User, status mtproto.TLObject) *mtproto.TL_user {
	tluser := &mtproto.TL_user{
		M_self:           mtproto.ToBool(true),
		M_contact:        mtproto.ToBool(true),
		M_mutual_contact: mtproto.ToBool(true),
		M_id:             int32(user.ID),
		M_access_hash:    user.AccessHash,
		M_first_name:     user.FirstName,
		M_last_name:      user.LastName,
		M_username:       user.UserName,
		M_phone:          user.Phone,
		M_status:         status,
	}

	return tluser
}

func makeUserStatusOnline() *mtproto.TL_userStatusOnline {
	now := time.Now().Unix()
	status := &mtproto.TL_userStatusOnline{
		M_expires: int32(now + 60),
	}
	return status
}

func makeTLUserByUser(selfid int64, user *User) mtproto.TLObject {
	if user == nil {
		return nil
	} else {
		var (
			status                 mtproto.TLObject
			photo                  mtproto.TLObject
			phone                  string
			contact, mutualContact bool
			isSelf                 = selfid == user.ID
		)

		if isSelf {
			status = makeUserStatusOnline()
			contact = true
			mutualContact = true
			phone = user.Phone
		} else {
			status = modelManager.GetUserStatus(user.ID)
			contact, mutualContact = modelManager.CheckContactAndMutualByUserID(selfid, user.ID)
			if contact {
				phone = user.Phone
			}
		}

		photoId := modelManager.GetDefaultUserPhotoID(user.ID)
		if photoId == 0 {
			photo = mtproto.New_TL_userProfilePhotoEmpty()
		} else {
			// photo = m.photoCallback.GetUserProfilePhoto(photoId)
			photo = mtproto.New_TL_userProfilePhotoEmpty() // 暂时不处理头像
		}

		data := &mtproto.TL_user{
			M_id:             int32(user.ID),
			M_self:           mtproto.ToBool(isSelf),
			M_contact:        mtproto.ToBool(contact),
			M_mutual_contact: mtproto.ToBool(mutualContact),
			M_access_hash:    user.AccessHash,
			M_first_name:     user.FirstName,
			M_last_name:      user.LastName,
			M_username:       user.UserName,
			M_phone:          phone,
			M_photo:          photo,
			M_status:         status,
		}

		return data
	}
}

func DialogList_to_TL_dialogList(dialogs []*UserDialog) []mtproto.TLObject {

	var tldialogs []mtproto.TLObject
	for _, d := range dialogs {

		tld := &mtproto.TL_dialog{
			M_pinned:                mtproto.ToBool(d.IsPinned == 1),
			M_peer:                  ToPeerByTypeAndID(d.PeerType, d.PeerID),
			M_top_message:           d.TopMessage,
			M_read_inbox_max_id:     d.ReadInboxMaxID,
			M_read_outbox_max_id:    d.ReadOutboxMaxID,
			M_unread_count:          d.UnreadCount,
			M_unread_mentions_count: d.UnreadMentionsCount,
		}

		if d.PeerType == PEER_CHANNEL {
			tld.M_pts = d.Pts
		}

		if d.DraftType == 2 {
			tldraft := &mtproto.TL_draftMessage{}
			err := UnserializeFromJson(d.DraftMessageData, &tldraft)
			if err != nil {
				Log.Errorf("DraftMessageData = %v, err = %v", FormatJson(d.DraftMessageData), err)
			} else {
				tld.M_draft = tldraft
			}
		}

		peerNotifySettings := &mtproto.TL_peerNotifySettings{
			M_show_previews: mtproto.ToBool(d.ShowPreviews == 1),
			M_silent:        mtproto.ToBool(d.Silent == 1),
			M_mute_until:    d.MuteUntil,
			M_sound:         d.Sound,
		}
		tld.M_notify_settings = peerNotifySettings

		tldialogs = append(tldialogs, tld)
	}

	return tldialogs
}

func Message_to_TL_message(message *Message) mtproto.TLObject {

	switch message.MessageType {
	case MESSAGE_TYPE_MESSAGE_EMPTY:
		return mtproto.New_TL_messageEmpty()
	case MESSAGE_TYPE_MESSAGE:
		m := &mtproto.TL_message{}
		err := UnserializeFromJson(message.MessageData, m)
		if err != nil {
			Log.Error(err)
			return nil
		} else {
			return m
		}
	case MESSAGE_TYPE_MESSAGE_SERVICE:
		m := &mtproto.TL_messageService{}
		err := UnserializeFromJson(message.MessageData, m)
		if err != nil {
			Log.Error(err)
			return nil
		} else {
			return m
		}

	default:
		Log.Error("unknown message type = %v", message.MessageType)
		return nil
	}
}

func MessageList_to_TL_messageList(messages []*Message) []mtproto.TLObject {
	if len(messages) == 0 {
		var ms []mtproto.TLObject
		return ms
	} else {
		var list []mtproto.TLObject
		for _, m := range messages {
			tlm := Message_to_TL_message(m)
			list = append(list, tlm)
		}
		return list
	}
}

func Chat_to_TL_chat(chat *Chat, selfid int64) mtproto.TLObject {

	tlchat := &mtproto.TL_chat{
		M_creator:        mtproto.ToBool(chat.CreatorUserID == selfid),
		M_id:             int32(chat.ID),
		M_title:          chat.Title,
		M_admins_enabled: mtproto.ToBool(chat.AdminsEnabled == 1),
		// Photo:             mtproto.NewTLChatPhotoEmpty().To_ChatPhoto(),
		M_participants_count: chat.ParticipantCount,
		M_date:               chat.Date,
		M_version:            chat.Version,
	}

	if chat.PhotoID == 0 {
		tlchat.M_photo = mtproto.New_TL_photoEmpty()
	} else {
		// chat.SetPhoto(this.cb.GetChatPhoto(this.chat.PhotoId))
		// 暂时不支持
		tlchat.M_photo = mtproto.New_TL_photoEmpty()
	}
	return tlchat
}

func ToPeerByTypeAndID(peerType int8, peerID int32) (peer mtproto.TLObject) {
	switch peerType {
	case PEER_USER:
		peer = &mtproto.TL_peerUser{
			M_user_id: peerID,
		}
	case PEER_CHAT:
		peer = &mtproto.TL_peerChat{
			M_chat_id: peerID,
		}
	case PEER_CHANNEL:
		peer = &mtproto.TL_peerChannel{
			M_channel_id: peerID,
		}
	default:
		panic(fmt.Sprintf("ToPeerByTypeAndID(%d, %d) error!", peerType, peerID))
	}
	return
}
