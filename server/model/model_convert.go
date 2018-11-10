package model

import (
	"fmt"

	. "github.com/rockin0098/meow/base/global"
	"github.com/rockin0098/meow/proto/mtproto"
)

const (
	PEER_EMPTY   = 0
	PEER_SELF    = 1
	PEER_USER    = 2
	PEER_CHAT    = 3
	PEER_CHANNEL = 4
	PEER_USERS   = 5
	PEER_CHATS   = 6
	PEER_ALL     = 7
	PEER_UNKNOWN = -1
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

func DialogList_to_TL_dialogList(dialogs []*UserDialog) []*mtproto.TL_dialog {

	var tldialogs []*mtproto.TL_dialog
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
