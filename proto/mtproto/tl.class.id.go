package mtproto

const (
	TL_LAYER_VERSION = "73"
)

const (
	TL_CLASS_UNKNOWN                                          int32 = 0
	TL_CLASS_resPQ                                            int32 = 85337187
	TL_CLASS_p_q_inner_data                                   int32 = -2083955988
	TL_CLASS_server_DH_params_fail                            int32 = 2043348061
	TL_CLASS_server_DH_params_ok                              int32 = -790100132
	TL_CLASS_server_DH_inner_data                             int32 = -1249309254
	TL_CLASS_client_DH_inner_data                             int32 = 1715713620
	TL_CLASS_dh_gen_ok                                        int32 = 1003222836
	TL_CLASS_dh_gen_retry                                     int32 = 1188831161
	TL_CLASS_dh_gen_fail                                      int32 = -1499615742
	TL_CLASS_destroy_auth_key_ok                              int32 = -161422892
	TL_CLASS_destroy_auth_key_none                            int32 = 178201177
	TL_CLASS_destroy_auth_key_fail                            int32 = -368010477
	TL_CLASS_req_pq                                           int32 = 1615239032
	TL_CLASS_req_DH_params                                    int32 = -686627650
	TL_CLASS_set_client_DH_params                             int32 = -184262881
	TL_CLASS_destroy_auth_key                                 int32 = -784117408
	TL_CLASS_msgs_ack                                         int32 = 1658238041
	TL_CLASS_bad_msg_notification                             int32 = -1477445615
	TL_CLASS_bad_server_salt                                  int32 = -307542917
	TL_CLASS_msgs_state_req                                   int32 = -630588590
	TL_CLASS_msgs_state_info                                  int32 = 81704317
	TL_CLASS_msgs_all_info                                    int32 = -1933520591
	TL_CLASS_msg_detailed_info                                int32 = 661470918
	TL_CLASS_msg_new_detailed_info                            int32 = -2137147681
	TL_CLASS_msg_resend_req                                   int32 = 2105940488
	TL_CLASS_rpc_error                                        int32 = 558156313
	TL_CLASS_rpc_answer_unknown                               int32 = 1579864942
	TL_CLASS_rpc_answer_dropped_running                       int32 = -847714938
	TL_CLASS_rpc_answer_dropped                               int32 = -1539647305
	TL_CLASS_future_salt                                      int32 = 155834844
	TL_CLASS_future_salts                                     int32 = -1370486635
	TL_CLASS_pong                                             int32 = 880243653
	TL_CLASS_destroy_session_ok                               int32 = -501201412
	TL_CLASS_destroy_session_none                             int32 = 1658015945
	TL_CLASS_new_session_created                              int32 = -1631450872
	TL_CLASS_http_wait                                        int32 = -1835453025
	TL_CLASS_ipPort                                           int32 = -734810765
	TL_CLASS_help_configSimple                                int32 = -644365371
	TL_CLASS_rpc_drop_answer                                  int32 = 1491380032
	TL_CLASS_get_future_salts                                 int32 = -1188971260
	TL_CLASS_ping                                             int32 = 2059302892
	TL_CLASS_ping_delay_disconnect                            int32 = -213746804
	TL_CLASS_destroy_session                                  int32 = -414113498
	TL_CLASS_contest_saveDeveloperInfo                        int32 = -1705021803
	TL_CLASS_boolFalse                                        int32 = -1132882121
	TL_CLASS_boolTrue                                         int32 = -1720552011
	TL_CLASS_true                                             int32 = 1072550713
	TL_CLASS_vector                                           int32 = 481674261
	TL_CLASS_error                                            int32 = -994444869
	TL_CLASS_null                                             int32 = 1450380236
	TL_CLASS_inputPeerEmpty                                   int32 = 2134579434
	TL_CLASS_inputPeerSelf                                    int32 = 2107670217
	TL_CLASS_inputPeerChat                                    int32 = 396093539
	TL_CLASS_inputPeerUser                                    int32 = 2072935910
	TL_CLASS_inputPeerChannel                                 int32 = 548253432
	TL_CLASS_inputUserEmpty                                   int32 = -1182234929
	TL_CLASS_inputUserSelf                                    int32 = -138301121
	TL_CLASS_inputUser                                        int32 = -668391402
	TL_CLASS_inputPhoneContact                                int32 = -208488460
	TL_CLASS_inputFile                                        int32 = -181407105
	TL_CLASS_inputFileBig                                     int32 = -95482955
	TL_CLASS_inputMediaEmpty                                  int32 = -1771768449
	TL_CLASS_inputMediaUploadedPhoto                          int32 = 792191537
	TL_CLASS_inputMediaPhoto                                  int32 = -2114308294
	TL_CLASS_inputMediaGeoPoint                               int32 = -104578748
	TL_CLASS_inputMediaContact                                int32 = -1494984313
	TL_CLASS_inputMediaUploadedDocument                       int32 = -476700163
	TL_CLASS_inputMediaDocument                               int32 = 1523279502
	TL_CLASS_inputMediaVenue                                  int32 = -1052959727
	TL_CLASS_inputMediaGifExternal                            int32 = 1212395773
	TL_CLASS_inputMediaPhotoExternal                          int32 = 153267905
	TL_CLASS_inputMediaDocumentExternal                       int32 = -1225309387
	TL_CLASS_inputMediaGame                                   int32 = -750828557
	TL_CLASS_inputMediaInvoice                                int32 = -186607933
	TL_CLASS_inputMediaGeoLive                                int32 = 2065305999
	TL_CLASS_inputChatPhotoEmpty                              int32 = 480546647
	TL_CLASS_inputChatUploadedPhoto                           int32 = -1837345356
	TL_CLASS_inputChatPhoto                                   int32 = -1991004873
	TL_CLASS_inputGeoPointEmpty                               int32 = -457104426
	TL_CLASS_inputGeoPoint                                    int32 = -206066487
	TL_CLASS_inputPhotoEmpty                                  int32 = 483901197
	TL_CLASS_inputPhoto                                       int32 = -74070332
	TL_CLASS_inputFileLocation                                int32 = 342061462
	TL_CLASS_inputEncryptedFileLocation                       int32 = -182231723
	TL_CLASS_inputDocumentFileLocation                        int32 = 1125058340
	TL_CLASS_inputAppEvent                                    int32 = 1996904104
	TL_CLASS_peerUser                                         int32 = -1649296275
	TL_CLASS_peerChat                                         int32 = -1160714821
	TL_CLASS_peerChannel                                      int32 = -1109531342
	TL_CLASS_storage_fileUnknown                              int32 = -1432995067
	TL_CLASS_storage_filePartial                              int32 = 1086091090
	TL_CLASS_storage_fileJpeg                                 int32 = 8322574
	TL_CLASS_storage_fileGif                                  int32 = -891180321
	TL_CLASS_storage_filePng                                  int32 = 172975040
	TL_CLASS_storage_filePdf                                  int32 = -1373745011
	TL_CLASS_storage_fileMp3                                  int32 = 1384777335
	TL_CLASS_storage_fileMov                                  int32 = 1258941372
	TL_CLASS_storage_fileMp4                                  int32 = -1278304028
	TL_CLASS_storage_fileWebp                                 int32 = 276907596
	TL_CLASS_fileLocationUnavailable                          int32 = 2086234950
	TL_CLASS_fileLocation                                     int32 = 1406570614
	TL_CLASS_userEmpty                                        int32 = 537022650
	TL_CLASS_user                                             int32 = 773059779
	TL_CLASS_userProfilePhotoEmpty                            int32 = 1326562017
	TL_CLASS_userProfilePhoto                                 int32 = -715532088
	TL_CLASS_userStatusEmpty                                  int32 = 164646985
	TL_CLASS_userStatusOnline                                 int32 = -306628279
	TL_CLASS_userStatusOffline                                int32 = 9203775
	TL_CLASS_userStatusRecently                               int32 = -496024847
	TL_CLASS_userStatusLastWeek                               int32 = 129960444
	TL_CLASS_userStatusLastMonth                              int32 = 2011940674
	TL_CLASS_chatEmpty                                        int32 = -1683826688
	TL_CLASS_chat                                             int32 = -652419756
	TL_CLASS_chatForbidden                                    int32 = 120753115
	TL_CLASS_channel                                          int32 = 1158377749
	TL_CLASS_channelForbidden                                 int32 = 681420594
	TL_CLASS_chatFull                                         int32 = 771925524
	TL_CLASS_channelFull                                      int32 = 1991201921
	TL_CLASS_chatParticipant                                  int32 = -925415106
	TL_CLASS_chatParticipantCreator                           int32 = -636267638
	TL_CLASS_chatParticipantAdmin                             int32 = -489233354
	TL_CLASS_chatParticipantsForbidden                        int32 = -57668565
	TL_CLASS_chatParticipants                                 int32 = 1061556205
	TL_CLASS_chatPhotoEmpty                                   int32 = 935395612
	TL_CLASS_chatPhoto                                        int32 = 1632839530
	TL_CLASS_messageEmpty                                     int32 = -2082087340
	TL_CLASS_message                                          int32 = 1157215293
	TL_CLASS_messageService                                   int32 = -1642487306
	TL_CLASS_messageMediaEmpty                                int32 = 1038967584
	TL_CLASS_messageMediaPhoto                                int32 = -1256047857
	TL_CLASS_messageMediaGeo                                  int32 = 1457575028
	TL_CLASS_messageMediaContact                              int32 = 1585262393
	TL_CLASS_messageMediaUnsupported                          int32 = -1618676578
	TL_CLASS_messageMediaDocument                             int32 = 2084836563
	TL_CLASS_messageMediaWebPage                              int32 = -1557277184
	TL_CLASS_messageMediaVenue                                int32 = 784356159
	TL_CLASS_messageMediaGame                                 int32 = -38694904
	TL_CLASS_messageMediaInvoice                              int32 = -2074799289
	TL_CLASS_messageMediaGeoLive                              int32 = 2084316681
	TL_CLASS_messageActionEmpty                               int32 = -1230047312
	TL_CLASS_messageActionChatCreate                          int32 = -1503425638
	TL_CLASS_messageActionChatEditTitle                       int32 = -1247687078
	TL_CLASS_messageActionChatEditPhoto                       int32 = 2144015272
	TL_CLASS_messageActionChatDeletePhoto                     int32 = -1780220945
	TL_CLASS_messageActionChatAddUser                         int32 = 1217033015
	TL_CLASS_messageActionChatDeleteUser                      int32 = -1297179892
	TL_CLASS_messageActionChatJoinedByLink                    int32 = -123931160
	TL_CLASS_messageActionChannelCreate                       int32 = -1781355374
	TL_CLASS_messageActionChatMigrateTo                       int32 = 1371385889
	TL_CLASS_messageActionChannelMigrateFrom                  int32 = -1336546578
	TL_CLASS_messageActionPinMessage                          int32 = -1799538451
	TL_CLASS_messageActionHistoryClear                        int32 = -1615153660
	TL_CLASS_messageActionGameScore                           int32 = -1834538890
	TL_CLASS_messageActionPaymentSentMe                       int32 = -1892568281
	TL_CLASS_messageActionPaymentSent                         int32 = 1080663248
	TL_CLASS_messageActionPhoneCall                           int32 = -2132731265
	TL_CLASS_messageActionScreenshotTaken                     int32 = 1200788123
	TL_CLASS_messageActionCustomAction                        int32 = -85549226
	TL_CLASS_dialog                                           int32 = -455150117
	TL_CLASS_photoEmpty                                       int32 = 590459437
	TL_CLASS_photo                                            int32 = -1836524247
	TL_CLASS_photoSizeEmpty                                   int32 = 236446268
	TL_CLASS_photoSize                                        int32 = 2009052699
	TL_CLASS_photoCachedSize                                  int32 = -374917894
	TL_CLASS_geoPointEmpty                                    int32 = 286776671
	TL_CLASS_geoPoint                                         int32 = 541710092
	TL_CLASS_auth_checkedPhone                                int32 = -2128698738
	TL_CLASS_auth_sentCode                                    int32 = 1577067778
	TL_CLASS_auth_authorization                               int32 = -855308010
	TL_CLASS_auth_exportedAuthorization                       int32 = -543777747
	TL_CLASS_inputNotifyPeer                                  int32 = -1195615476
	TL_CLASS_inputNotifyUsers                                 int32 = 423314455
	TL_CLASS_inputNotifyChats                                 int32 = 1251338318
	TL_CLASS_inputNotifyAll                                   int32 = -1540769658
	TL_CLASS_inputPeerNotifyEventsEmpty                       int32 = -265263912
	TL_CLASS_inputPeerNotifyEventsAll                         int32 = -395694988
	TL_CLASS_inputPeerNotifySettings                          int32 = 949182130
	TL_CLASS_peerNotifyEventsEmpty                            int32 = -1378534221
	TL_CLASS_peerNotifyEventsAll                              int32 = 1830677896
	TL_CLASS_peerNotifySettingsEmpty                          int32 = 1889961234
	TL_CLASS_peerNotifySettings                               int32 = -1697798976
	TL_CLASS_peerSettings                                     int32 = -2122045747
	TL_CLASS_wallPaper                                        int32 = -860866985
	TL_CLASS_wallPaperSolid                                   int32 = 1662091044
	TL_CLASS_inputReportReasonSpam                            int32 = 1490799288
	TL_CLASS_inputReportReasonViolence                        int32 = 505595789
	TL_CLASS_inputReportReasonPornography                     int32 = 777640226
	TL_CLASS_inputReportReasonOther                           int32 = -512463606
	TL_CLASS_userFull                                         int32 = 253890367
	TL_CLASS_contact                                          int32 = -116274796
	TL_CLASS_importedContact                                  int32 = -805141448
	TL_CLASS_contactBlocked                                   int32 = 1444661369
	TL_CLASS_contactStatus                                    int32 = -748155807
	TL_CLASS_contacts_link                                    int32 = 986597452
	TL_CLASS_contacts_contactsNotModified                     int32 = -1219778094
	TL_CLASS_contacts_contacts                                int32 = -353862078
	TL_CLASS_contacts_importedContacts                        int32 = 2010127419
	TL_CLASS_contacts_blocked                                 int32 = 471043349
	TL_CLASS_contacts_blockedSlice                            int32 = -1878523231
	TL_CLASS_messages_dialogs                                 int32 = 364538944
	TL_CLASS_messages_dialogsSlice                            int32 = 1910543603
	TL_CLASS_messages_messages                                int32 = -1938715001
	TL_CLASS_messages_messagesSlice                           int32 = 189033187
	TL_CLASS_messages_channelMessages                         int32 = -1725551049
	TL_CLASS_messages_messagesNotModified                     int32 = 1951620897
	TL_CLASS_messages_chats                                   int32 = 1694474197
	TL_CLASS_messages_chatsSlice                              int32 = -1663561404
	TL_CLASS_messages_chatFull                                int32 = -438840932
	TL_CLASS_messages_affectedHistory                         int32 = -1269012015
	TL_CLASS_inputMessagesFilterEmpty                         int32 = 1474492012
	TL_CLASS_inputMessagesFilterPhotos                        int32 = -1777752804
	TL_CLASS_inputMessagesFilterVideo                         int32 = -1614803355
	TL_CLASS_inputMessagesFilterPhotoVideo                    int32 = 1458172132
	TL_CLASS_inputMessagesFilterDocument                      int32 = -1629621880
	TL_CLASS_inputMessagesFilterUrl                           int32 = 2129714567
	TL_CLASS_inputMessagesFilterGif                           int32 = -3644025
	TL_CLASS_inputMessagesFilterVoice                         int32 = 1358283666
	TL_CLASS_inputMessagesFilterMusic                         int32 = 928101534
	TL_CLASS_inputMessagesFilterChatPhotos                    int32 = 975236280
	TL_CLASS_inputMessagesFilterPhoneCalls                    int32 = -2134272152
	TL_CLASS_inputMessagesFilterRoundVoice                    int32 = 2054952868
	TL_CLASS_inputMessagesFilterRoundVideo                    int32 = -1253451181
	TL_CLASS_inputMessagesFilterMyMentions                    int32 = -1040652646
	TL_CLASS_inputMessagesFilterGeo                           int32 = -419271411
	TL_CLASS_inputMessagesFilterContacts                      int32 = -530392189
	TL_CLASS_updateNewMessage                                 int32 = 522914557
	TL_CLASS_updateMessageID                                  int32 = 1318109142
	TL_CLASS_updateDeleteMessages                             int32 = -1576161051
	TL_CLASS_updateUserTyping                                 int32 = 1548249383
	TL_CLASS_updateChatUserTyping                             int32 = -1704596961
	TL_CLASS_updateChatParticipants                           int32 = 125178264
	TL_CLASS_updateUserStatus                                 int32 = 469489699
	TL_CLASS_updateUserName                                   int32 = -1489818765
	TL_CLASS_updateUserPhoto                                  int32 = -1791935732
	TL_CLASS_updateContactRegistered                          int32 = 628472761
	TL_CLASS_updateContactLink                                int32 = -1657903163
	TL_CLASS_updateNewEncryptedMessage                        int32 = 314359194
	TL_CLASS_updateEncryptedChatTyping                        int32 = 386986326
	TL_CLASS_updateEncryption                                 int32 = -1264392051
	TL_CLASS_updateEncryptedMessagesRead                      int32 = 956179895
	TL_CLASS_updateChatParticipantAdd                         int32 = -364179876
	TL_CLASS_updateChatParticipantDelete                      int32 = 1851755554
	TL_CLASS_updateDcOptions                                  int32 = -1906403213
	TL_CLASS_updateUserBlocked                                int32 = -2131957734
	TL_CLASS_updateNotifySettings                             int32 = -1094555409
	TL_CLASS_updateServiceNotification                        int32 = -337352679
	TL_CLASS_updatePrivacy                                    int32 = -298113238
	TL_CLASS_updateUserPhone                                  int32 = 314130811
	TL_CLASS_updateReadHistoryInbox                           int32 = -1721631396
	TL_CLASS_updateReadHistoryOutbox                          int32 = 791617983
	TL_CLASS_updateWebPage                                    int32 = 2139689491
	TL_CLASS_updateReadMessagesContents                       int32 = 1757493555
	TL_CLASS_updateChannelTooLong                             int32 = -352032773
	TL_CLASS_updateChannel                                    int32 = -1227598250
	TL_CLASS_updateNewChannelMessage                          int32 = 1656358105
	TL_CLASS_updateReadChannelInbox                           int32 = 1108669311
	TL_CLASS_updateDeleteChannelMessages                      int32 = -1015733815
	TL_CLASS_updateChannelMessageViews                        int32 = -1734268085
	TL_CLASS_updateChatAdmins                                 int32 = 1855224129
	TL_CLASS_updateChatParticipantAdmin                       int32 = -1232070311
	TL_CLASS_updateNewStickerSet                              int32 = 1753886890
	TL_CLASS_updateStickerSetsOrder                           int32 = 196268545
	TL_CLASS_updateStickerSets                                int32 = 1135492588
	TL_CLASS_updateSavedGifs                                  int32 = -1821035490
	TL_CLASS_updateBotInlineQuery                             int32 = 1417832080
	TL_CLASS_updateBotInlineSend                              int32 = 239663460
	TL_CLASS_updateEditChannelMessage                         int32 = 457133559
	TL_CLASS_updateChannelPinnedMessage                       int32 = -1738988427
	TL_CLASS_updateBotCallbackQuery                           int32 = -415938591
	TL_CLASS_updateEditMessage                                int32 = -469536605
	TL_CLASS_updateInlineBotCallbackQuery                     int32 = -103646630
	TL_CLASS_updateReadChannelOutbox                          int32 = 634833351
	TL_CLASS_updateDraftMessage                               int32 = -299124375
	TL_CLASS_updateReadFeaturedStickers                       int32 = 1461528386
	TL_CLASS_updateRecentStickers                             int32 = -1706939360
	TL_CLASS_updateConfig                                     int32 = -1574314746
	TL_CLASS_updatePtsChanged                                 int32 = 861169551
	TL_CLASS_updateChannelWebPage                             int32 = 1081547008
	TL_CLASS_updateDialogPinned                               int32 = -686710068
	TL_CLASS_updatePinnedDialogs                              int32 = -657787251
	TL_CLASS_updateBotWebhookJSON                             int32 = -2095595325
	TL_CLASS_updateBotWebhookJSONQuery                        int32 = -1684914010
	TL_CLASS_updateBotShippingQuery                           int32 = -523384512
	TL_CLASS_updateBotPrecheckoutQuery                        int32 = 1563376297
	TL_CLASS_updatePhoneCall                                  int32 = -1425052898
	TL_CLASS_updateLangPackTooLong                            int32 = 281165899
	TL_CLASS_updateLangPack                                   int32 = 1442983757
	TL_CLASS_updateFavedStickers                              int32 = -451831443
	TL_CLASS_updateChannelReadMessagesContents                int32 = -1987495099
	TL_CLASS_updateContactsReset                              int32 = 1887741886
	TL_CLASS_updateChannelAvailableMessages                   int32 = 1893427255
	TL_CLASS_updates_state                                    int32 = -1519637954
	TL_CLASS_updates_differenceEmpty                          int32 = 1567990072
	TL_CLASS_updates_difference                               int32 = 16030880
	TL_CLASS_updates_differenceSlice                          int32 = -1459938943
	TL_CLASS_updates_differenceTooLong                        int32 = 1258196845
	TL_CLASS_updatesTooLong                                   int32 = -484987010
	TL_CLASS_updateShortMessage                               int32 = -1857044719
	TL_CLASS_updateShortChatMessage                           int32 = 377562760
	TL_CLASS_updateShort                                      int32 = 2027216577
	TL_CLASS_updatesCombined                                  int32 = 1918567619
	TL_CLASS_updates                                          int32 = 1957577280
	TL_CLASS_updateShortSentMessage                           int32 = 301019932
	TL_CLASS_photos_photos                                    int32 = -1916114267
	TL_CLASS_photos_photosSlice                               int32 = 352657236
	TL_CLASS_photos_photo                                     int32 = 539045032
	TL_CLASS_upload_file                                      int32 = 157948117
	TL_CLASS_upload_fileCdnRedirect                           int32 = -363659686
	TL_CLASS_dcOption                                         int32 = 98092748
	TL_CLASS_config                                           int32 = -1669068444
	TL_CLASS_nearestDc                                        int32 = -1910892683
	TL_CLASS_help_appUpdate                                   int32 = -1987579119
	TL_CLASS_help_noAppUpdate                                 int32 = -1000708810
	TL_CLASS_help_inviteText                                  int32 = 415997816
	TL_CLASS_encryptedChatEmpty                               int32 = -1417756512
	TL_CLASS_encryptedChatWaiting                             int32 = 1006044124
	TL_CLASS_encryptedChatRequested                           int32 = -931638658
	TL_CLASS_encryptedChat                                    int32 = -94974410
	TL_CLASS_encryptedChatDiscarded                           int32 = 332848423
	TL_CLASS_inputEncryptedChat                               int32 = -247351839
	TL_CLASS_encryptedFileEmpty                               int32 = -1038136962
	TL_CLASS_encryptedFile                                    int32 = 1248893260
	TL_CLASS_inputEncryptedFileEmpty                          int32 = 406307684
	TL_CLASS_inputEncryptedFileUploaded                       int32 = 1690108678
	TL_CLASS_inputEncryptedFile                               int32 = 1511503333
	TL_CLASS_inputEncryptedFileBigUploaded                    int32 = 767652808
	TL_CLASS_encryptedMessage                                 int32 = -317144808
	TL_CLASS_encryptedMessageService                          int32 = 594758406
	TL_CLASS_messages_dhConfigNotModified                     int32 = -1058912715
	TL_CLASS_messages_dhConfig                                int32 = 740433629
	TL_CLASS_messages_sentEncryptedMessage                    int32 = 1443858741
	TL_CLASS_messages_sentEncryptedFile                       int32 = -1802240206
	TL_CLASS_inputDocumentEmpty                               int32 = 1928391342
	TL_CLASS_inputDocument                                    int32 = 410618194
	TL_CLASS_documentEmpty                                    int32 = 922273905
	TL_CLASS_document                                         int32 = -2027738169
	TL_CLASS_help_support                                     int32 = 398898678
	TL_CLASS_notifyPeer                                       int32 = -1613493288
	TL_CLASS_notifyUsers                                      int32 = -1261946036
	TL_CLASS_notifyChats                                      int32 = -1073230141
	TL_CLASS_notifyAll                                        int32 = 1959820384
	TL_CLASS_sendMessageTypingAction                          int32 = 381645902
	TL_CLASS_sendMessageCancelAction                          int32 = -44119819
	TL_CLASS_sendMessageRecordVideoAction                     int32 = -1584933265
	TL_CLASS_sendMessageUploadVideoAction                     int32 = -378127636
	TL_CLASS_sendMessageRecordAudioAction                     int32 = -718310409
	TL_CLASS_sendMessageUploadAudioAction                     int32 = -212740181
	TL_CLASS_sendMessageUploadPhotoAction                     int32 = -774682074
	TL_CLASS_sendMessageUploadDocumentAction                  int32 = -1441998364
	TL_CLASS_sendMessageGeoLocationAction                     int32 = 393186209
	TL_CLASS_sendMessageChooseContactAction                   int32 = 1653390447
	TL_CLASS_sendMessageGamePlayAction                        int32 = -580219064
	TL_CLASS_sendMessageRecordRoundAction                     int32 = -1997373508
	TL_CLASS_sendMessageUploadRoundAction                     int32 = 608050278
	TL_CLASS_contacts_found                                   int32 = 446822276
	TL_CLASS_inputPrivacyKeyStatusTimestamp                   int32 = 1335282456
	TL_CLASS_inputPrivacyKeyChatInvite                        int32 = -1107622874
	TL_CLASS_inputPrivacyKeyPhoneCall                         int32 = -88417185
	TL_CLASS_privacyKeyStatusTimestamp                        int32 = -1137792208
	TL_CLASS_privacyKeyChatInvite                             int32 = 1343122938
	TL_CLASS_privacyKeyPhoneCall                              int32 = 1030105979
	TL_CLASS_inputPrivacyValueAllowContacts                   int32 = 218751099
	TL_CLASS_inputPrivacyValueAllowAll                        int32 = 407582158
	TL_CLASS_inputPrivacyValueAllowUsers                      int32 = 320652927
	TL_CLASS_inputPrivacyValueDisallowContacts                int32 = 195371015
	TL_CLASS_inputPrivacyValueDisallowAll                     int32 = -697604407
	TL_CLASS_inputPrivacyValueDisallowUsers                   int32 = -1877932953
	TL_CLASS_privacyValueAllowContacts                        int32 = -123988
	TL_CLASS_privacyValueAllowAll                             int32 = 1698855810
	TL_CLASS_privacyValueAllowUsers                           int32 = 1297858060
	TL_CLASS_privacyValueDisallowContacts                     int32 = -125240806
	TL_CLASS_privacyValueDisallowAll                          int32 = -1955338397
	TL_CLASS_privacyValueDisallowUsers                        int32 = 209668535
	TL_CLASS_account_privacyRules                             int32 = 1430961007
	TL_CLASS_accountDaysTTL                                   int32 = -1194283041
	TL_CLASS_documentAttributeImageSize                       int32 = 1815593308
	TL_CLASS_documentAttributeAnimated                        int32 = 297109817
	TL_CLASS_documentAttributeSticker                         int32 = 1662637586
	TL_CLASS_documentAttributeVideo                           int32 = 250621158
	TL_CLASS_documentAttributeAudio                           int32 = -1739392570
	TL_CLASS_documentAttributeFilename                        int32 = 358154344
	TL_CLASS_documentAttributeHasStickers                     int32 = -1744710921
	TL_CLASS_messages_stickersNotModified                     int32 = -244016606
	TL_CLASS_messages_stickers                                int32 = -1970352846
	TL_CLASS_stickerPack                                      int32 = 313694676
	TL_CLASS_messages_allStickersNotModified                  int32 = -395967805
	TL_CLASS_messages_allStickers                             int32 = -302170017
	TL_CLASS_disabledFeature                                  int32 = -1369215196
	TL_CLASS_messages_affectedMessages                        int32 = -2066640507
	TL_CLASS_contactLinkUnknown                               int32 = 1599050311
	TL_CLASS_contactLinkNone                                  int32 = -17968211
	TL_CLASS_contactLinkHasPhone                              int32 = 646922073
	TL_CLASS_contactLinkContact                               int32 = -721239344
	TL_CLASS_webPageEmpty                                     int32 = -350980120
	TL_CLASS_webPagePending                                   int32 = -981018084
	TL_CLASS_webPage                                          int32 = 1594340540
	TL_CLASS_webPageNotModified                               int32 = -2054908813
	TL_CLASS_authorization                                    int32 = 2079516406
	TL_CLASS_account_authorizations                           int32 = 307276766
	TL_CLASS_account_noPassword                               int32 = -1764049896
	TL_CLASS_account_password                                 int32 = 2081952796
	TL_CLASS_account_passwordSettings                         int32 = -1212732749
	TL_CLASS_account_passwordInputSettings                    int32 = -2037289493
	TL_CLASS_auth_passwordRecovery                            int32 = 326715557
	TL_CLASS_receivedNotifyMessage                            int32 = -1551583367
	TL_CLASS_chatInviteEmpty                                  int32 = 1776236393
	TL_CLASS_chatInviteExported                               int32 = -64092740
	TL_CLASS_chatInviteAlready                                int32 = 1516793212
	TL_CLASS_chatInvite                                       int32 = -613092008
	TL_CLASS_inputStickerSetEmpty                             int32 = -4838507
	TL_CLASS_inputStickerSetID                                int32 = -1645763991
	TL_CLASS_inputStickerSetShortName                         int32 = -2044933984
	TL_CLASS_stickerSet                                       int32 = -852477119
	TL_CLASS_messages_stickerSet                              int32 = -1240849242
	TL_CLASS_botCommand                                       int32 = -1032140601
	TL_CLASS_botInfo                                          int32 = -1729618630
	TL_CLASS_keyboardButton                                   int32 = -1560655744
	TL_CLASS_keyboardButtonUrl                                int32 = 629866245
	TL_CLASS_keyboardButtonCallback                           int32 = 1748655686
	TL_CLASS_keyboardButtonRequestPhone                       int32 = -1318425559
	TL_CLASS_keyboardButtonRequestGeoLocation                 int32 = -59151553
	TL_CLASS_keyboardButtonSwitchInline                       int32 = 90744648
	TL_CLASS_keyboardButtonGame                               int32 = 1358175439
	TL_CLASS_keyboardButtonBuy                                int32 = -1344716869
	TL_CLASS_keyboardButtonRow                                int32 = 2002815875
	TL_CLASS_replyKeyboardHide                                int32 = -1606526075
	TL_CLASS_replyKeyboardForceReply                          int32 = -200242528
	TL_CLASS_replyKeyboardMarkup                              int32 = 889353612
	TL_CLASS_replyInlineMarkup                                int32 = 1218642516
	TL_CLASS_messageEntityUnknown                             int32 = -1148011883
	TL_CLASS_messageEntityMention                             int32 = -100378723
	TL_CLASS_messageEntityHashtag                             int32 = 1868782349
	TL_CLASS_messageEntityBotCommand                          int32 = 1827637959
	TL_CLASS_messageEntityUrl                                 int32 = 1859134776
	TL_CLASS_messageEntityEmail                               int32 = 1692693954
	TL_CLASS_messageEntityBold                                int32 = -1117713463
	TL_CLASS_messageEntityItalic                              int32 = -2106619040
	TL_CLASS_messageEntityCode                                int32 = 681706865
	TL_CLASS_messageEntityPre                                 int32 = 1938967520
	TL_CLASS_messageEntityTextUrl                             int32 = 1990644519
	TL_CLASS_messageEntityMentionName                         int32 = 892193368
	TL_CLASS_inputMessageEntityMentionName                    int32 = 546203849
	TL_CLASS_inputChannelEmpty                                int32 = -292807034
	TL_CLASS_inputChannel                                     int32 = -1343524562
	TL_CLASS_contacts_resolvedPeer                            int32 = 2131196633
	TL_CLASS_messageRange                                     int32 = 182649427
	TL_CLASS_updates_channelDifferenceEmpty                   int32 = 1041346555
	TL_CLASS_updates_channelDifferenceTooLong                 int32 = 1788705589
	TL_CLASS_updates_channelDifference                        int32 = 543450958
	TL_CLASS_channelMessagesFilterEmpty                       int32 = -1798033689
	TL_CLASS_channelMessagesFilter                            int32 = -847783593
	TL_CLASS_channelParticipant                               int32 = 367766557
	TL_CLASS_channelParticipantSelf                           int32 = -1557620115
	TL_CLASS_channelParticipantCreator                        int32 = -471670279
	TL_CLASS_channelParticipantAdmin                          int32 = -1473271656
	TL_CLASS_channelParticipantBanned                         int32 = 573315206
	TL_CLASS_channelParticipantsRecent                        int32 = -566281095
	TL_CLASS_channelParticipantsAdmins                        int32 = -1268741783
	TL_CLASS_channelParticipantsKicked                        int32 = -1548400251
	TL_CLASS_channelParticipantsBots                          int32 = -1328445861
	TL_CLASS_channelParticipantsBanned                        int32 = 338142689
	TL_CLASS_channelParticipantsSearch                        int32 = 106343499
	TL_CLASS_channels_channelParticipants                     int32 = -177282392
	TL_CLASS_channels_channelParticipantsNotModified          int32 = -266911767
	TL_CLASS_channels_channelParticipant                      int32 = -791039645
	TL_CLASS_help_termsOfService                              int32 = -236044656
	TL_CLASS_foundGif                                         int32 = 372165663
	TL_CLASS_foundGifCached                                   int32 = -1670052855
	TL_CLASS_messages_foundGifs                               int32 = 1158290442
	TL_CLASS_messages_savedGifsNotModified                    int32 = -402498398
	TL_CLASS_messages_savedGifs                               int32 = 772213157
	TL_CLASS_inputBotInlineMessageMediaAuto                   int32 = 691006739
	TL_CLASS_inputBotInlineMessageText                        int32 = 1036876423
	TL_CLASS_inputBotInlineMessageMediaGeo                    int32 = -1045340827
	TL_CLASS_inputBotInlineMessageMediaVenue                  int32 = -1431327288
	TL_CLASS_inputBotInlineMessageMediaContact                int32 = 766443943
	TL_CLASS_inputBotInlineMessageGame                        int32 = 1262639204
	TL_CLASS_inputBotInlineResult                             int32 = 750510426
	TL_CLASS_inputBotInlineResultPhoto                        int32 = -1462213465
	TL_CLASS_inputBotInlineResultDocument                     int32 = -459324
	TL_CLASS_inputBotInlineResultGame                         int32 = 1336154098
	TL_CLASS_botInlineMessageMediaAuto                        int32 = 175419739
	TL_CLASS_botInlineMessageText                             int32 = -1937807902
	TL_CLASS_botInlineMessageMediaGeo                         int32 = -1222451611
	TL_CLASS_botInlineMessageMediaVenue                       int32 = 1130767150
	TL_CLASS_botInlineMessageMediaContact                     int32 = 904770772
	TL_CLASS_botInlineResult                                  int32 = -1679053127
	TL_CLASS_botInlineMediaResult                             int32 = 400266251
	TL_CLASS_messages_botResults                              int32 = -1803769784
	TL_CLASS_exportedMessageLink                              int32 = 524838915
	TL_CLASS_messageFwdHeader                                 int32 = 1436466797
	TL_CLASS_auth_codeTypeSms                                 int32 = 1923290508
	TL_CLASS_auth_codeTypeCall                                int32 = 1948046307
	TL_CLASS_auth_codeTypeFlashCall                           int32 = 577556219
	TL_CLASS_auth_sentCodeTypeApp                             int32 = 1035688326
	TL_CLASS_auth_sentCodeTypeSms                             int32 = -1073693790
	TL_CLASS_auth_sentCodeTypeCall                            int32 = 1398007207
	TL_CLASS_auth_sentCodeTypeFlashCall                       int32 = -1425815847
	TL_CLASS_messages_botCallbackAnswer                       int32 = 911761060
	TL_CLASS_messages_messageEditData                         int32 = 649453030
	TL_CLASS_inputBotInlineMessageID                          int32 = -1995686519
	TL_CLASS_inlineBotSwitchPM                                int32 = 1008755359
	TL_CLASS_messages_peerDialogs                             int32 = 863093588
	TL_CLASS_topPeer                                          int32 = -305282981
	TL_CLASS_topPeerCategoryBotsPM                            int32 = -1419371685
	TL_CLASS_topPeerCategoryBotsInline                        int32 = 344356834
	TL_CLASS_topPeerCategoryCorrespondents                    int32 = 104314861
	TL_CLASS_topPeerCategoryGroups                            int32 = -1122524854
	TL_CLASS_topPeerCategoryChannels                          int32 = 371037736
	TL_CLASS_topPeerCategoryPhoneCalls                        int32 = 511092620
	TL_CLASS_topPeerCategoryPeers                             int32 = -75283823
	TL_CLASS_contacts_topPeersNotModified                     int32 = -567906571
	TL_CLASS_contacts_topPeers                                int32 = 1891070632
	TL_CLASS_draftMessageEmpty                                int32 = -1169445179
	TL_CLASS_draftMessage                                     int32 = -40996577
	TL_CLASS_messages_featuredStickersNotModified             int32 = 82699215
	TL_CLASS_messages_featuredStickers                        int32 = -123893531
	TL_CLASS_messages_recentStickersNotModified               int32 = 186120336
	TL_CLASS_messages_recentStickers                          int32 = 1558317424
	TL_CLASS_messages_archivedStickers                        int32 = 1338747336
	TL_CLASS_messages_stickerSetInstallResultSuccess          int32 = 946083368
	TL_CLASS_messages_stickerSetInstallResultArchive          int32 = 904138920
	TL_CLASS_stickerSetCovered                                int32 = 1678812626
	TL_CLASS_stickerSetMultiCovered                           int32 = 872932635
	TL_CLASS_maskCoords                                       int32 = -1361650766
	TL_CLASS_inputStickeredMediaPhoto                         int32 = 1251549527
	TL_CLASS_inputStickeredMediaDocument                      int32 = 70813275
	TL_CLASS_game                                             int32 = -1107729093
	TL_CLASS_inputGameID                                      int32 = 53231223
	TL_CLASS_inputGameShortName                               int32 = -1020139510
	TL_CLASS_highScore                                        int32 = 1493171408
	TL_CLASS_messages_highScores                              int32 = -1707344487
	TL_CLASS_textEmpty                                        int32 = -599948721
	TL_CLASS_textPlain                                        int32 = 1950782688
	TL_CLASS_textBold                                         int32 = 1730456516
	TL_CLASS_textItalic                                       int32 = -653089380
	TL_CLASS_textUnderline                                    int32 = -1054465340
	TL_CLASS_textStrike                                       int32 = -1678197867
	TL_CLASS_textFixed                                        int32 = 1816074681
	TL_CLASS_textUrl                                          int32 = 1009288385
	TL_CLASS_textEmail                                        int32 = -564523562
	TL_CLASS_textConcat                                       int32 = 2120376535
	TL_CLASS_pageBlockUnsupported                             int32 = 324435594
	TL_CLASS_pageBlockTitle                                   int32 = 1890305021
	TL_CLASS_pageBlockSubtitle                                int32 = -1879401953
	TL_CLASS_pageBlockAuthorDate                              int32 = -1162877472
	TL_CLASS_pageBlockHeader                                  int32 = -1076861716
	TL_CLASS_pageBlockSubheader                               int32 = -248793375
	TL_CLASS_pageBlockParagraph                               int32 = 1182402406
	TL_CLASS_pageBlockPreformatted                            int32 = -1066346178
	TL_CLASS_pageBlockFooter                                  int32 = 1216809369
	TL_CLASS_pageBlockDivider                                 int32 = -618614392
	TL_CLASS_pageBlockAnchor                                  int32 = -837994576
	TL_CLASS_pageBlockList                                    int32 = 978896884
	TL_CLASS_pageBlockBlockquote                              int32 = 641563686
	TL_CLASS_pageBlockPullquote                               int32 = 1329878739
	TL_CLASS_pageBlockPhoto                                   int32 = -372860542
	TL_CLASS_pageBlockVideo                                   int32 = -640214938
	TL_CLASS_pageBlockCover                                   int32 = 972174080
	TL_CLASS_pageBlockEmbed                                   int32 = -840826671
	TL_CLASS_pageBlockEmbedPost                               int32 = 690781161
	TL_CLASS_pageBlockCollage                                 int32 = 145955919
	TL_CLASS_pageBlockSlideshow                               int32 = 319588707
	TL_CLASS_pageBlockChannel                                 int32 = -283684427
	TL_CLASS_pageBlockAudio                                   int32 = 834148991
	TL_CLASS_pagePart                                         int32 = -1908433218
	TL_CLASS_pageFull                                         int32 = 1433323434
	TL_CLASS_phoneCallDiscardReasonMissed                     int32 = -2048646399
	TL_CLASS_phoneCallDiscardReasonDisconnect                 int32 = -527056480
	TL_CLASS_phoneCallDiscardReasonHangup                     int32 = 1471006352
	TL_CLASS_phoneCallDiscardReasonBusy                       int32 = -84416311
	TL_CLASS_dataJSON                                         int32 = 2104790276
	TL_CLASS_labeledPrice                                     int32 = -886477832
	TL_CLASS_invoice                                          int32 = -1022713000
	TL_CLASS_paymentCharge                                    int32 = -368917890
	TL_CLASS_postAddress                                      int32 = 512535275
	TL_CLASS_paymentRequestedInfo                             int32 = -1868808300
	TL_CLASS_paymentSavedCredentialsCard                      int32 = -842892769
	TL_CLASS_webDocument                                      int32 = -971322408
	TL_CLASS_inputWebDocument                                 int32 = -1678949555
	TL_CLASS_inputWebFileLocation                             int32 = -1036396922
	TL_CLASS_upload_webFile                                   int32 = 568808380
	TL_CLASS_payments_paymentForm                             int32 = 1062645411
	TL_CLASS_payments_validatedRequestedInfo                  int32 = -784000893
	TL_CLASS_payments_paymentResult                           int32 = 1314881805
	TL_CLASS_payments_paymentVerficationNeeded                int32 = 1800845601
	TL_CLASS_payments_paymentReceipt                          int32 = 1342771681
	TL_CLASS_payments_savedInfo                               int32 = -74456004
	TL_CLASS_inputPaymentCredentialsSaved                     int32 = -1056001329
	TL_CLASS_inputPaymentCredentials                          int32 = 873977640
	TL_CLASS_inputPaymentCredentialsApplePay                  int32 = 178373535
	TL_CLASS_inputPaymentCredentialsAndroidPay                int32 = 2035705766
	TL_CLASS_account_tmpPassword                              int32 = -614138572
	TL_CLASS_shippingOption                                   int32 = -1239335713
	TL_CLASS_inputStickerSetItem                              int32 = -6249322
	TL_CLASS_inputPhoneCall                                   int32 = 506920429
	TL_CLASS_phoneCallEmpty                                   int32 = 1399245077
	TL_CLASS_phoneCallWaiting                                 int32 = 462375633
	TL_CLASS_phoneCallRequested                               int32 = -2089411356
	TL_CLASS_phoneCallAccepted                                int32 = 1828732223
	TL_CLASS_phoneCall                                        int32 = -1660057
	TL_CLASS_phoneCallDiscarded                               int32 = 1355435489
	TL_CLASS_phoneConnection                                  int32 = -1655957568
	TL_CLASS_phoneCallProtocol                                int32 = -1564789301
	TL_CLASS_phone_phoneCall                                  int32 = -326966976
	TL_CLASS_upload_cdnFileReuploadNeeded                     int32 = -290921362
	TL_CLASS_upload_cdnFile                                   int32 = -1449145777
	TL_CLASS_cdnPublicKey                                     int32 = -914167110
	TL_CLASS_cdnConfig                                        int32 = 1462101002
	TL_CLASS_langPackString                                   int32 = -892239370
	TL_CLASS_langPackStringPluralized                         int32 = 1816636575
	TL_CLASS_langPackStringDeleted                            int32 = 695856818
	TL_CLASS_langPackDifference                               int32 = -209337866
	TL_CLASS_langPackLanguage                                 int32 = 292985073
	TL_CLASS_channelAdminRights                               int32 = 1568467877
	TL_CLASS_channelBannedRights                              int32 = 1489977929
	TL_CLASS_channelAdminLogEventActionChangeTitle            int32 = -421545947
	TL_CLASS_channelAdminLogEventActionChangeAbout            int32 = 1427671598
	TL_CLASS_channelAdminLogEventActionChangeUsername         int32 = 1783299128
	TL_CLASS_channelAdminLogEventActionChangePhoto            int32 = -1204857405
	TL_CLASS_channelAdminLogEventActionToggleInvites          int32 = 460916654
	TL_CLASS_channelAdminLogEventActionToggleSignatures       int32 = 648939889
	TL_CLASS_channelAdminLogEventActionUpdatePinned           int32 = -370660328
	TL_CLASS_channelAdminLogEventActionEditMessage            int32 = 1889215493
	TL_CLASS_channelAdminLogEventActionDeleteMessage          int32 = 1121994683
	TL_CLASS_channelAdminLogEventActionParticipantJoin        int32 = 405815507
	TL_CLASS_channelAdminLogEventActionParticipantLeave       int32 = -124291086
	TL_CLASS_channelAdminLogEventActionParticipantInvite      int32 = -484690728
	TL_CLASS_channelAdminLogEventActionParticipantToggleBan   int32 = -422036098
	TL_CLASS_channelAdminLogEventActionParticipantToggleAdmin int32 = -714643696
	TL_CLASS_channelAdminLogEventActionChangeStickerSet       int32 = -1312568665
	TL_CLASS_channelAdminLogEventActionTogglePreHistoryHidden int32 = 1599903217
	TL_CLASS_channelAdminLogEvent                             int32 = 995769920
	TL_CLASS_channels_adminLogResults                         int32 = -309659827
	TL_CLASS_channelAdminLogEventsFilter                      int32 = -368018716
	TL_CLASS_popularContact                                   int32 = 1558266229
	TL_CLASS_cdnFileHash                                      int32 = 2012136335
	TL_CLASS_messages_favedStickersNotModified                int32 = -1634752813
	TL_CLASS_messages_favedStickers                           int32 = -209768682
	TL_CLASS_recentMeUrlUnknown                               int32 = 1189204285
	TL_CLASS_recentMeUrlUser                                  int32 = -1917045962
	TL_CLASS_recentMeUrlChat                                  int32 = -1608834311
	TL_CLASS_recentMeUrlChatInvite                            int32 = -347535331
	TL_CLASS_recentMeUrlStickerSet                            int32 = -1140172836
	TL_CLASS_help_recentMeUrls                                int32 = 235081943
	TL_CLASS_inputSingleMedia                                 int32 = 1588230153
	TL_CLASS_invokeAfterMsg                                   int32 = -878758099
	TL_CLASS_invokeAfterMsgs                                  int32 = 1036301552
	TL_CLASS_initConnection                                   int32 = -951575130
	TL_CLASS_invokeWithLayer                                  int32 = -627372787
	TL_CLASS_invokeWithoutUpdates                             int32 = -1080796745
	TL_CLASS_auth_checkPhone                                  int32 = 1877286395
	TL_CLASS_auth_sendCode                                    int32 = -2035355412
	TL_CLASS_auth_signUp                                      int32 = 453408308
	TL_CLASS_auth_signIn                                      int32 = -1126886015
	TL_CLASS_auth_logOut                                      int32 = 1461180992
	TL_CLASS_auth_resetAuthorizations                         int32 = -1616179942
	TL_CLASS_auth_sendInvites                                 int32 = 1998331287
	TL_CLASS_auth_exportAuthorization                         int32 = -440401971
	TL_CLASS_auth_importAuthorization                         int32 = -470837741
	TL_CLASS_auth_bindTempAuthKey                             int32 = -841733627
	TL_CLASS_auth_importBotAuthorization                      int32 = 1738800940
	TL_CLASS_auth_checkPassword                               int32 = 174260510
	TL_CLASS_auth_requestPasswordRecovery                     int32 = -661144474
	TL_CLASS_auth_recoverPassword                             int32 = 1319464594
	TL_CLASS_auth_resendCode                                  int32 = 1056025023
	TL_CLASS_auth_cancelCode                                  int32 = 520357240
	TL_CLASS_auth_dropTempAuthKeys                            int32 = -1907842680
	TL_CLASS_account_registerDevice                           int32 = 1669245048
	TL_CLASS_account_unregisterDevice                         int32 = 1707432768
	TL_CLASS_account_updateNotifySettings                     int32 = -2067899501
	TL_CLASS_account_getNotifySettings                        int32 = 313765169
	TL_CLASS_account_resetNotifySettings                      int32 = -612493497
	TL_CLASS_account_updateProfile                            int32 = 2018596725
	TL_CLASS_account_updateStatus                             int32 = 1713919532
	TL_CLASS_account_getWallPapers                            int32 = -1068696894
	TL_CLASS_account_reportPeer                               int32 = -1374118561
	TL_CLASS_account_checkUsername                            int32 = 655677548
	TL_CLASS_account_updateUsername                           int32 = 1040964988
	TL_CLASS_account_getPrivacy                               int32 = -623130288
	TL_CLASS_account_setPrivacy                               int32 = -906486552
	TL_CLASS_account_deleteAccount                            int32 = 1099779595
	TL_CLASS_account_getAccountTTL                            int32 = 150761757
	TL_CLASS_account_setAccountTTL                            int32 = 608323678
	TL_CLASS_account_sendChangePhoneCode                      int32 = 149257707
	TL_CLASS_account_changePhone                              int32 = 1891839707
	TL_CLASS_account_updateDeviceLocked                       int32 = 954152242
	TL_CLASS_account_getAuthorizations                        int32 = -484392616
	TL_CLASS_account_resetAuthorization                       int32 = -545786948
	TL_CLASS_account_getPassword                              int32 = 1418342645
	TL_CLASS_account_getPasswordSettings                      int32 = -1131605573
	TL_CLASS_account_updatePasswordSettings                   int32 = -92517498
	TL_CLASS_account_sendConfirmPhoneCode                     int32 = 353818557
	TL_CLASS_account_confirmPhone                             int32 = 1596029123
	TL_CLASS_account_getTmpPassword                           int32 = 1250046590
	TL_CLASS_users_getUsers                                   int32 = 227648840
	TL_CLASS_users_getFullUser                                int32 = -902781519
	TL_CLASS_contacts_getStatuses                             int32 = -995929106
	TL_CLASS_contacts_getContacts                             int32 = -1071414113
	TL_CLASS_contacts_importContacts                          int32 = 746589157
	TL_CLASS_contacts_deleteContact                           int32 = -1902823612
	TL_CLASS_contacts_deleteContacts                          int32 = 1504393374
	TL_CLASS_contacts_block                                   int32 = 858475004
	TL_CLASS_contacts_unblock                                 int32 = -448724803
	TL_CLASS_contacts_getBlocked                              int32 = -176409329
	TL_CLASS_contacts_exportCard                              int32 = -2065352905
	TL_CLASS_contacts_importCard                              int32 = 1340184318
	TL_CLASS_contacts_search                                  int32 = 301470424
	TL_CLASS_contacts_resolveUsername                         int32 = -113456221
	TL_CLASS_contacts_getTopPeers                             int32 = -728224331
	TL_CLASS_contacts_resetTopPeerRating                      int32 = 451113900
	TL_CLASS_contacts_resetSaved                              int32 = -2020263951
	TL_CLASS_messages_getMessages                             int32 = 1109588596
	TL_CLASS_messages_getDialogs                              int32 = 421243333
	TL_CLASS_messages_getHistory                              int32 = -591691168
	TL_CLASS_messages_search                                  int32 = 60726944
	TL_CLASS_messages_readHistory                             int32 = 238054714
	TL_CLASS_messages_deleteHistory                           int32 = 469850889
	TL_CLASS_messages_deleteMessages                          int32 = -443640366
	TL_CLASS_messages_receivedMessages                        int32 = 94983360
	TL_CLASS_messages_setTyping                               int32 = -1551737264
	TL_CLASS_messages_sendMessage                             int32 = -91733382
	TL_CLASS_messages_sendMedia                               int32 = -923703407
	TL_CLASS_messages_forwardMessages                         int32 = 1888354709
	TL_CLASS_messages_reportSpam                              int32 = -820669733
	TL_CLASS_messages_hideReportSpam                          int32 = -1460572005
	TL_CLASS_messages_getPeerSettings                         int32 = 913498268
	TL_CLASS_messages_getChats                                int32 = 1013621127
	TL_CLASS_messages_getFullChat                             int32 = 998448230
	TL_CLASS_messages_editChatTitle                           int32 = -599447467
	TL_CLASS_messages_editChatPhoto                           int32 = -900957736
	TL_CLASS_messages_addChatUser                             int32 = -106911223
	TL_CLASS_messages_deleteChatUser                          int32 = -530505962
	TL_CLASS_messages_createChat                              int32 = 164303470
	TL_CLASS_messages_forwardMessage                          int32 = 865483769
	TL_CLASS_messages_getDhConfig                             int32 = 651135312
	TL_CLASS_messages_requestEncryption                       int32 = -162681021
	TL_CLASS_messages_acceptEncryption                        int32 = 1035731989
	TL_CLASS_messages_discardEncryption                       int32 = -304536635
	TL_CLASS_messages_setEncryptedTyping                      int32 = 2031374829
	TL_CLASS_messages_readEncryptedHistory                    int32 = 2135648522
	TL_CLASS_messages_sendEncrypted                           int32 = -1451792525
	TL_CLASS_messages_sendEncryptedFile                       int32 = -1701831834
	TL_CLASS_messages_sendEncryptedService                    int32 = 852769188
	TL_CLASS_messages_receivedQueue                           int32 = 1436924774
	TL_CLASS_messages_reportEncryptedSpam                     int32 = 1259113487
	TL_CLASS_messages_readMessageContents                     int32 = 916930423
	TL_CLASS_messages_getAllStickers                          int32 = 479598769
	TL_CLASS_messages_getWebPagePreview                       int32 = 623001124
	TL_CLASS_messages_exportChatInvite                        int32 = 2106086025
	TL_CLASS_messages_checkChatInvite                         int32 = 1051570619
	TL_CLASS_messages_importChatInvite                        int32 = 1817183516
	TL_CLASS_messages_getStickerSet                           int32 = 639215886
	TL_CLASS_messages_installStickerSet                       int32 = -946871200
	TL_CLASS_messages_uninstallStickerSet                     int32 = -110209570
	TL_CLASS_messages_startBot                                int32 = -421563528
	TL_CLASS_messages_getMessagesViews                        int32 = -993483427
	TL_CLASS_messages_toggleChatAdmins                        int32 = -326379039
	TL_CLASS_messages_editChatAdmin                           int32 = -1444503762
	TL_CLASS_messages_migrateChat                             int32 = 363051235
	TL_CLASS_messages_searchGlobal                            int32 = -1640190800
	TL_CLASS_messages_reorderStickerSets                      int32 = 2016638777
	TL_CLASS_messages_getDocumentByHash                       int32 = 864953444
	TL_CLASS_messages_searchGifs                              int32 = -1080395925
	TL_CLASS_messages_getSavedGifs                            int32 = -2084618926
	TL_CLASS_messages_saveGif                                 int32 = 846868683
	TL_CLASS_messages_getInlineBotResults                     int32 = 1364105629
	TL_CLASS_messages_setInlineBotResults                     int32 = -346119674
	TL_CLASS_messages_sendInlineBotResult                     int32 = -1318189314
	TL_CLASS_messages_getMessageEditData                      int32 = -39416522
	TL_CLASS_messages_editMessage                             int32 = 97630429
	TL_CLASS_messages_editInlineBotMessage                    int32 = -1327463869
	TL_CLASS_messages_getBotCallbackAnswer                    int32 = -2130010132
	TL_CLASS_messages_setBotCallbackAnswer                    int32 = -712043766
	TL_CLASS_messages_getPeerDialogs                          int32 = 764901049
	TL_CLASS_messages_saveDraft                               int32 = -1137057461
	TL_CLASS_messages_getAllDrafts                            int32 = 1782549861
	TL_CLASS_messages_getFeaturedStickers                     int32 = 766298703
	TL_CLASS_messages_readFeaturedStickers                    int32 = 1527873830
	TL_CLASS_messages_getRecentStickers                       int32 = 1587647177
	TL_CLASS_messages_saveRecentSticker                       int32 = 958863608
	TL_CLASS_messages_clearRecentStickers                     int32 = -1986437075
	TL_CLASS_messages_getArchivedStickers                     int32 = 1475442322
	TL_CLASS_messages_getMaskStickers                         int32 = 1706608543
	TL_CLASS_messages_getAttachedStickers                     int32 = -866424884
	TL_CLASS_messages_setGameScore                            int32 = -1896289088
	TL_CLASS_messages_setInlineGameScore                      int32 = 363700068
	TL_CLASS_messages_getGameHighScores                       int32 = -400399203
	TL_CLASS_messages_getInlineGameHighScores                 int32 = 258170395
	TL_CLASS_messages_getCommonChats                          int32 = 218777796
	TL_CLASS_messages_getAllChats                             int32 = -341307408
	TL_CLASS_messages_getWebPage                              int32 = 852135825
	TL_CLASS_messages_toggleDialogPin                         int32 = 847887978
	TL_CLASS_messages_reorderPinnedDialogs                    int32 = -1784678844
	TL_CLASS_messages_getPinnedDialogs                        int32 = -497756594
	TL_CLASS_messages_setBotShippingResults                   int32 = -436833542
	TL_CLASS_messages_setBotPrecheckoutResults                int32 = 163765653
	TL_CLASS_messages_uploadMedia                             int32 = 1369162417
	TL_CLASS_messages_sendScreenshotNotification              int32 = -914493408
	TL_CLASS_messages_getFavedStickers                        int32 = 567151374
	TL_CLASS_messages_faveSticker                             int32 = -1174420133
	TL_CLASS_messages_getUnreadMentions                       int32 = 1180140658
	TL_CLASS_messages_readMentions                            int32 = 251759059
	TL_CLASS_messages_getRecentLocations                      int32 = 613691874
	TL_CLASS_messages_sendMultiMedia                          int32 = 546656559
	TL_CLASS_messages_uploadEncryptedFile                     int32 = 1347929239
	TL_CLASS_updates_getState                                 int32 = -304838614
	TL_CLASS_updates_getDifference                            int32 = 630429265
	TL_CLASS_updates_getChannelDifference                     int32 = 51854712
	TL_CLASS_photos_updateProfilePhoto                        int32 = -256159406
	TL_CLASS_photos_uploadProfilePhoto                        int32 = 1328726168
	TL_CLASS_photos_deletePhotos                              int32 = -2016444625
	TL_CLASS_photos_getUserPhotos                             int32 = -1848823128
	TL_CLASS_upload_saveFilePart                              int32 = -1291540959
	TL_CLASS_upload_getFile                                   int32 = -475607115
	TL_CLASS_upload_saveBigFilePart                           int32 = -562337987
	TL_CLASS_upload_getWebFile                                int32 = 619086221
	TL_CLASS_upload_getCdnFile                                int32 = 536919235
	TL_CLASS_upload_reuploadCdnFile                           int32 = 452533257
	TL_CLASS_upload_getCdnFileHashes                          int32 = -149567365
	TL_CLASS_help_getConfig                                   int32 = -990308245
	TL_CLASS_help_getNearestDc                                int32 = 531836966
	TL_CLASS_help_getAppUpdate                                int32 = -1372724842
	TL_CLASS_help_saveAppLog                                  int32 = 1862465352
	TL_CLASS_help_getInviteText                               int32 = 1295590211
	TL_CLASS_help_getSupport                                  int32 = -1663104819
	TL_CLASS_help_getAppChangelog                             int32 = -1877938321
	TL_CLASS_help_getTermsOfService                           int32 = 889286899
	TL_CLASS_help_setBotUpdatesStatus                         int32 = -333262899
	TL_CLASS_help_getCdnConfig                                int32 = 1375900482
	TL_CLASS_help_getRecentMeUrls                             int32 = 1036054804
	TL_CLASS_channels_readHistory                             int32 = -871347913
	TL_CLASS_channels_deleteMessages                          int32 = -2067661490
	TL_CLASS_channels_deleteUserHistory                       int32 = -787622117
	TL_CLASS_channels_reportSpam                              int32 = -32999408
	TL_CLASS_channels_getMessages                             int32 = -1814580409
	TL_CLASS_channels_getParticipants                         int32 = 306054633
	TL_CLASS_channels_getParticipant                          int32 = 1416484774
	TL_CLASS_channels_getChannels                             int32 = 176122811
	TL_CLASS_channels_getFullChannel                          int32 = 141781513
	TL_CLASS_channels_createChannel                           int32 = -192332417
	TL_CLASS_channels_editAbout                               int32 = 333610782
	TL_CLASS_channels_editAdmin                               int32 = 548962836
	TL_CLASS_channels_editTitle                               int32 = 1450044624
	TL_CLASS_channels_editPhoto                               int32 = -248621111
	TL_CLASS_channels_checkUsername                           int32 = 283557164
	TL_CLASS_channels_updateUsername                          int32 = 890549214
	TL_CLASS_channels_joinChannel                             int32 = 615851205
	TL_CLASS_channels_leaveChannel                            int32 = -130635115
	TL_CLASS_channels_inviteToChannel                         int32 = 429865580
	TL_CLASS_channels_exportInvite                            int32 = -950663035
	TL_CLASS_channels_deleteChannel                           int32 = -1072619549
	TL_CLASS_channels_toggleInvites                           int32 = 1231065863
	TL_CLASS_channels_exportMessageLink                       int32 = -934882771
	TL_CLASS_channels_toggleSignatures                        int32 = 527021574
	TL_CLASS_channels_updatePinnedMessage                     int32 = -1490162350
	TL_CLASS_channels_getAdminedPublicChannels                int32 = -1920105769
	TL_CLASS_channels_editBanned                              int32 = -1076292147
	TL_CLASS_channels_getAdminLog                             int32 = 870184064
	TL_CLASS_channels_setStickers                             int32 = -359881479
	TL_CLASS_channels_readMessageContents                     int32 = -357180360
	TL_CLASS_channels_deleteHistory                           int32 = -1355375294
	TL_CLASS_channels_togglePreHistoryHidden                  int32 = -356796084
	TL_CLASS_bots_sendCustomRequest                           int32 = -1440257555
	TL_CLASS_bots_answerWebhookJSONQuery                      int32 = -434028723
	TL_CLASS_payments_getPaymentForm                          int32 = -1712285883
	TL_CLASS_payments_getPaymentReceipt                       int32 = -1601001088
	TL_CLASS_payments_validateRequestedInfo                   int32 = 1997180532
	TL_CLASS_payments_sendPaymentForm                         int32 = 730364339
	TL_CLASS_payments_getSavedInfo                            int32 = 578650699
	TL_CLASS_payments_clearSavedInfo                          int32 = -667062079
	TL_CLASS_stickers_createStickerSet                        int32 = -1680314774
	TL_CLASS_stickers_removeStickerFromSet                    int32 = -143257775
	TL_CLASS_stickers_changeStickerPosition                   int32 = -4795190
	TL_CLASS_stickers_addStickerToSet                         int32 = -2041315650
	TL_CLASS_phone_getCallConfig                              int32 = 1430593449
	TL_CLASS_phone_requestCall                                int32 = 1536537556
	TL_CLASS_phone_acceptCall                                 int32 = 1003664544
	TL_CLASS_phone_confirmCall                                int32 = 788404002
	TL_CLASS_phone_receivedCall                               int32 = 399855457
	TL_CLASS_phone_discardCall                                int32 = 2027164582
	TL_CLASS_phone_setCallRating                              int32 = 475228724
	TL_CLASS_phone_saveCallDebug                              int32 = 662363518
	TL_CLASS_langpack_getLangPack                             int32 = -1699363442
	TL_CLASS_langpack_getStrings                              int32 = 773776152
	TL_CLASS_langpack_getDifference                           int32 = 187583869
	TL_CLASS_langpack_getLanguages                            int32 = -2146445955
)

var TL_CLASS_NAME = map[int32]string{
	85337187:    "TL_CLASS_resPQ",
	-2083955988: "TL_CLASS_p_q_inner_data",
	2043348061:  "TL_CLASS_server_DH_params_fail",
	-790100132:  "TL_CLASS_server_DH_params_ok",
	-1249309254: "TL_CLASS_server_DH_inner_data",
	1715713620:  "TL_CLASS_client_DH_inner_data",
	1003222836:  "TL_CLASS_dh_gen_ok",
	1188831161:  "TL_CLASS_dh_gen_retry",
	-1499615742: "TL_CLASS_dh_gen_fail",
	-161422892:  "TL_CLASS_destroy_auth_key_ok",
	178201177:   "TL_CLASS_destroy_auth_key_none",
	-368010477:  "TL_CLASS_destroy_auth_key_fail",
	1615239032:  "TL_CLASS_req_pq",
	-686627650:  "TL_CLASS_req_DH_params",
	-184262881:  "TL_CLASS_set_client_DH_params",
	-784117408:  "TL_CLASS_destroy_auth_key",
	1658238041:  "TL_CLASS_msgs_ack",
	-1477445615: "TL_CLASS_bad_msg_notification",
	-307542917:  "TL_CLASS_bad_server_salt",
	-630588590:  "TL_CLASS_msgs_state_req",
	81704317:    "TL_CLASS_msgs_state_info",
	-1933520591: "TL_CLASS_msgs_all_info",
	661470918:   "TL_CLASS_msg_detailed_info",
	-2137147681: "TL_CLASS_msg_new_detailed_info",
	2105940488:  "TL_CLASS_msg_resend_req",
	558156313:   "TL_CLASS_rpc_error",
	1579864942:  "TL_CLASS_rpc_answer_unknown",
	-847714938:  "TL_CLASS_rpc_answer_dropped_running",
	-1539647305: "TL_CLASS_rpc_answer_dropped",
	155834844:   "TL_CLASS_future_salt",
	-1370486635: "TL_CLASS_future_salts",
	880243653:   "TL_CLASS_pong",
	-501201412:  "TL_CLASS_destroy_session_ok",
	1658015945:  "TL_CLASS_destroy_session_none",
	-1631450872: "TL_CLASS_new_session_created",
	-1835453025: "TL_CLASS_http_wait",
	-734810765:  "TL_CLASS_ipPort",
	-644365371:  "TL_CLASS_help_configSimple",
	1491380032:  "TL_CLASS_rpc_drop_answer",
	-1188971260: "TL_CLASS_get_future_salts",
	2059302892:  "TL_CLASS_ping",
	-213746804:  "TL_CLASS_ping_delay_disconnect",
	-414113498:  "TL_CLASS_destroy_session",
	-1705021803: "TL_CLASS_contest_saveDeveloperInfo",
	-1132882121: "TL_CLASS_boolFalse",
	-1720552011: "TL_CLASS_boolTrue",
	1072550713:  "TL_CLASS_true",
	481674261:   "TL_CLASS_vector",
	-994444869:  "TL_CLASS_error",
	1450380236:  "TL_CLASS_null",
	2134579434:  "TL_CLASS_inputPeerEmpty",
	2107670217:  "TL_CLASS_inputPeerSelf",
	396093539:   "TL_CLASS_inputPeerChat",
	2072935910:  "TL_CLASS_inputPeerUser",
	548253432:   "TL_CLASS_inputPeerChannel",
	-1182234929: "TL_CLASS_inputUserEmpty",
	-138301121:  "TL_CLASS_inputUserSelf",
	-668391402:  "TL_CLASS_inputUser",
	-208488460:  "TL_CLASS_inputPhoneContact",
	-181407105:  "TL_CLASS_inputFile",
	-95482955:   "TL_CLASS_inputFileBig",
	-1771768449: "TL_CLASS_inputMediaEmpty",
	792191537:   "TL_CLASS_inputMediaUploadedPhoto",
	-2114308294: "TL_CLASS_inputMediaPhoto",
	-104578748:  "TL_CLASS_inputMediaGeoPoint",
	-1494984313: "TL_CLASS_inputMediaContact",
	-476700163:  "TL_CLASS_inputMediaUploadedDocument",
	1523279502:  "TL_CLASS_inputMediaDocument",
	-1052959727: "TL_CLASS_inputMediaVenue",
	1212395773:  "TL_CLASS_inputMediaGifExternal",
	153267905:   "TL_CLASS_inputMediaPhotoExternal",
	-1225309387: "TL_CLASS_inputMediaDocumentExternal",
	-750828557:  "TL_CLASS_inputMediaGame",
	-186607933:  "TL_CLASS_inputMediaInvoice",
	2065305999:  "TL_CLASS_inputMediaGeoLive",
	480546647:   "TL_CLASS_inputChatPhotoEmpty",
	-1837345356: "TL_CLASS_inputChatUploadedPhoto",
	-1991004873: "TL_CLASS_inputChatPhoto",
	-457104426:  "TL_CLASS_inputGeoPointEmpty",
	-206066487:  "TL_CLASS_inputGeoPoint",
	483901197:   "TL_CLASS_inputPhotoEmpty",
	-74070332:   "TL_CLASS_inputPhoto",
	342061462:   "TL_CLASS_inputFileLocation",
	-182231723:  "TL_CLASS_inputEncryptedFileLocation",
	1125058340:  "TL_CLASS_inputDocumentFileLocation",
	1996904104:  "TL_CLASS_inputAppEvent",
	-1649296275: "TL_CLASS_peerUser",
	-1160714821: "TL_CLASS_peerChat",
	-1109531342: "TL_CLASS_peerChannel",
	-1432995067: "TL_CLASS_storage_fileUnknown",
	1086091090:  "TL_CLASS_storage_filePartial",
	8322574:     "TL_CLASS_storage_fileJpeg",
	-891180321:  "TL_CLASS_storage_fileGif",
	172975040:   "TL_CLASS_storage_filePng",
	-1373745011: "TL_CLASS_storage_filePdf",
	1384777335:  "TL_CLASS_storage_fileMp3",
	1258941372:  "TL_CLASS_storage_fileMov",
	-1278304028: "TL_CLASS_storage_fileMp4",
	276907596:   "TL_CLASS_storage_fileWebp",
	2086234950:  "TL_CLASS_fileLocationUnavailable",
	1406570614:  "TL_CLASS_fileLocation",
	537022650:   "TL_CLASS_userEmpty",
	773059779:   "TL_CLASS_user",
	1326562017:  "TL_CLASS_userProfilePhotoEmpty",
	-715532088:  "TL_CLASS_userProfilePhoto",
	164646985:   "TL_CLASS_userStatusEmpty",
	-306628279:  "TL_CLASS_userStatusOnline",
	9203775:     "TL_CLASS_userStatusOffline",
	-496024847:  "TL_CLASS_userStatusRecently",
	129960444:   "TL_CLASS_userStatusLastWeek",
	2011940674:  "TL_CLASS_userStatusLastMonth",
	-1683826688: "TL_CLASS_chatEmpty",
	-652419756:  "TL_CLASS_chat",
	120753115:   "TL_CLASS_chatForbidden",
	1158377749:  "TL_CLASS_channel",
	681420594:   "TL_CLASS_channelForbidden",
	771925524:   "TL_CLASS_chatFull",
	1991201921:  "TL_CLASS_channelFull",
	-925415106:  "TL_CLASS_chatParticipant",
	-636267638:  "TL_CLASS_chatParticipantCreator",
	-489233354:  "TL_CLASS_chatParticipantAdmin",
	-57668565:   "TL_CLASS_chatParticipantsForbidden",
	1061556205:  "TL_CLASS_chatParticipants",
	935395612:   "TL_CLASS_chatPhotoEmpty",
	1632839530:  "TL_CLASS_chatPhoto",
	-2082087340: "TL_CLASS_messageEmpty",
	1157215293:  "TL_CLASS_message",
	-1642487306: "TL_CLASS_messageService",
	1038967584:  "TL_CLASS_messageMediaEmpty",
	-1256047857: "TL_CLASS_messageMediaPhoto",
	1457575028:  "TL_CLASS_messageMediaGeo",
	1585262393:  "TL_CLASS_messageMediaContact",
	-1618676578: "TL_CLASS_messageMediaUnsupported",
	2084836563:  "TL_CLASS_messageMediaDocument",
	-1557277184: "TL_CLASS_messageMediaWebPage",
	784356159:   "TL_CLASS_messageMediaVenue",
	-38694904:   "TL_CLASS_messageMediaGame",
	-2074799289: "TL_CLASS_messageMediaInvoice",
	2084316681:  "TL_CLASS_messageMediaGeoLive",
	-1230047312: "TL_CLASS_messageActionEmpty",
	-1503425638: "TL_CLASS_messageActionChatCreate",
	-1247687078: "TL_CLASS_messageActionChatEditTitle",
	2144015272:  "TL_CLASS_messageActionChatEditPhoto",
	-1780220945: "TL_CLASS_messageActionChatDeletePhoto",
	1217033015:  "TL_CLASS_messageActionChatAddUser",
	-1297179892: "TL_CLASS_messageActionChatDeleteUser",
	-123931160:  "TL_CLASS_messageActionChatJoinedByLink",
	-1781355374: "TL_CLASS_messageActionChannelCreate",
	1371385889:  "TL_CLASS_messageActionChatMigrateTo",
	-1336546578: "TL_CLASS_messageActionChannelMigrateFrom",
	-1799538451: "TL_CLASS_messageActionPinMessage",
	-1615153660: "TL_CLASS_messageActionHistoryClear",
	-1834538890: "TL_CLASS_messageActionGameScore",
	-1892568281: "TL_CLASS_messageActionPaymentSentMe",
	1080663248:  "TL_CLASS_messageActionPaymentSent",
	-2132731265: "TL_CLASS_messageActionPhoneCall",
	1200788123:  "TL_CLASS_messageActionScreenshotTaken",
	-85549226:   "TL_CLASS_messageActionCustomAction",
	-455150117:  "TL_CLASS_dialog",
	590459437:   "TL_CLASS_photoEmpty",
	-1836524247: "TL_CLASS_photo",
	236446268:   "TL_CLASS_photoSizeEmpty",
	2009052699:  "TL_CLASS_photoSize",
	-374917894:  "TL_CLASS_photoCachedSize",
	286776671:   "TL_CLASS_geoPointEmpty",
	541710092:   "TL_CLASS_geoPoint",
	-2128698738: "TL_CLASS_auth_checkedPhone",
	1577067778:  "TL_CLASS_auth_sentCode",
	-855308010:  "TL_CLASS_auth_authorization",
	-543777747:  "TL_CLASS_auth_exportedAuthorization",
	-1195615476: "TL_CLASS_inputNotifyPeer",
	423314455:   "TL_CLASS_inputNotifyUsers",
	1251338318:  "TL_CLASS_inputNotifyChats",
	-1540769658: "TL_CLASS_inputNotifyAll",
	-265263912:  "TL_CLASS_inputPeerNotifyEventsEmpty",
	-395694988:  "TL_CLASS_inputPeerNotifyEventsAll",
	949182130:   "TL_CLASS_inputPeerNotifySettings",
	-1378534221: "TL_CLASS_peerNotifyEventsEmpty",
	1830677896:  "TL_CLASS_peerNotifyEventsAll",
	1889961234:  "TL_CLASS_peerNotifySettingsEmpty",
	-1697798976: "TL_CLASS_peerNotifySettings",
	-2122045747: "TL_CLASS_peerSettings",
	-860866985:  "TL_CLASS_wallPaper",
	1662091044:  "TL_CLASS_wallPaperSolid",
	1490799288:  "TL_CLASS_inputReportReasonSpam",
	505595789:   "TL_CLASS_inputReportReasonViolence",
	777640226:   "TL_CLASS_inputReportReasonPornography",
	-512463606:  "TL_CLASS_inputReportReasonOther",
	253890367:   "TL_CLASS_userFull",
	-116274796:  "TL_CLASS_contact",
	-805141448:  "TL_CLASS_importedContact",
	1444661369:  "TL_CLASS_contactBlocked",
	-748155807:  "TL_CLASS_contactStatus",
	986597452:   "TL_CLASS_contacts_link",
	-1219778094: "TL_CLASS_contacts_contactsNotModified",
	-353862078:  "TL_CLASS_contacts_contacts",
	2010127419:  "TL_CLASS_contacts_importedContacts",
	471043349:   "TL_CLASS_contacts_blocked",
	-1878523231: "TL_CLASS_contacts_blockedSlice",
	364538944:   "TL_CLASS_messages_dialogs",
	1910543603:  "TL_CLASS_messages_dialogsSlice",
	-1938715001: "TL_CLASS_messages_messages",
	189033187:   "TL_CLASS_messages_messagesSlice",
	-1725551049: "TL_CLASS_messages_channelMessages",
	1951620897:  "TL_CLASS_messages_messagesNotModified",
	1694474197:  "TL_CLASS_messages_chats",
	-1663561404: "TL_CLASS_messages_chatsSlice",
	-438840932:  "TL_CLASS_messages_chatFull",
	-1269012015: "TL_CLASS_messages_affectedHistory",
	1474492012:  "TL_CLASS_inputMessagesFilterEmpty",
	-1777752804: "TL_CLASS_inputMessagesFilterPhotos",
	-1614803355: "TL_CLASS_inputMessagesFilterVideo",
	1458172132:  "TL_CLASS_inputMessagesFilterPhotoVideo",
	-1629621880: "TL_CLASS_inputMessagesFilterDocument",
	2129714567:  "TL_CLASS_inputMessagesFilterUrl",
	-3644025:    "TL_CLASS_inputMessagesFilterGif",
	1358283666:  "TL_CLASS_inputMessagesFilterVoice",
	928101534:   "TL_CLASS_inputMessagesFilterMusic",
	975236280:   "TL_CLASS_inputMessagesFilterChatPhotos",
	-2134272152: "TL_CLASS_inputMessagesFilterPhoneCalls",
	2054952868:  "TL_CLASS_inputMessagesFilterRoundVoice",
	-1253451181: "TL_CLASS_inputMessagesFilterRoundVideo",
	-1040652646: "TL_CLASS_inputMessagesFilterMyMentions",
	-419271411:  "TL_CLASS_inputMessagesFilterGeo",
	-530392189:  "TL_CLASS_inputMessagesFilterContacts",
	522914557:   "TL_CLASS_updateNewMessage",
	1318109142:  "TL_CLASS_updateMessageID",
	-1576161051: "TL_CLASS_updateDeleteMessages",
	1548249383:  "TL_CLASS_updateUserTyping",
	-1704596961: "TL_CLASS_updateChatUserTyping",
	125178264:   "TL_CLASS_updateChatParticipants",
	469489699:   "TL_CLASS_updateUserStatus",
	-1489818765: "TL_CLASS_updateUserName",
	-1791935732: "TL_CLASS_updateUserPhoto",
	628472761:   "TL_CLASS_updateContactRegistered",
	-1657903163: "TL_CLASS_updateContactLink",
	314359194:   "TL_CLASS_updateNewEncryptedMessage",
	386986326:   "TL_CLASS_updateEncryptedChatTyping",
	-1264392051: "TL_CLASS_updateEncryption",
	956179895:   "TL_CLASS_updateEncryptedMessagesRead",
	-364179876:  "TL_CLASS_updateChatParticipantAdd",
	1851755554:  "TL_CLASS_updateChatParticipantDelete",
	-1906403213: "TL_CLASS_updateDcOptions",
	-2131957734: "TL_CLASS_updateUserBlocked",
	-1094555409: "TL_CLASS_updateNotifySettings",
	-337352679:  "TL_CLASS_updateServiceNotification",
	-298113238:  "TL_CLASS_updatePrivacy",
	314130811:   "TL_CLASS_updateUserPhone",
	-1721631396: "TL_CLASS_updateReadHistoryInbox",
	791617983:   "TL_CLASS_updateReadHistoryOutbox",
	2139689491:  "TL_CLASS_updateWebPage",
	1757493555:  "TL_CLASS_updateReadMessagesContents",
	-352032773:  "TL_CLASS_updateChannelTooLong",
	-1227598250: "TL_CLASS_updateChannel",
	1656358105:  "TL_CLASS_updateNewChannelMessage",
	1108669311:  "TL_CLASS_updateReadChannelInbox",
	-1015733815: "TL_CLASS_updateDeleteChannelMessages",
	-1734268085: "TL_CLASS_updateChannelMessageViews",
	1855224129:  "TL_CLASS_updateChatAdmins",
	-1232070311: "TL_CLASS_updateChatParticipantAdmin",
	1753886890:  "TL_CLASS_updateNewStickerSet",
	196268545:   "TL_CLASS_updateStickerSetsOrder",
	1135492588:  "TL_CLASS_updateStickerSets",
	-1821035490: "TL_CLASS_updateSavedGifs",
	1417832080:  "TL_CLASS_updateBotInlineQuery",
	239663460:   "TL_CLASS_updateBotInlineSend",
	457133559:   "TL_CLASS_updateEditChannelMessage",
	-1738988427: "TL_CLASS_updateChannelPinnedMessage",
	-415938591:  "TL_CLASS_updateBotCallbackQuery",
	-469536605:  "TL_CLASS_updateEditMessage",
	-103646630:  "TL_CLASS_updateInlineBotCallbackQuery",
	634833351:   "TL_CLASS_updateReadChannelOutbox",
	-299124375:  "TL_CLASS_updateDraftMessage",
	1461528386:  "TL_CLASS_updateReadFeaturedStickers",
	-1706939360: "TL_CLASS_updateRecentStickers",
	-1574314746: "TL_CLASS_updateConfig",
	861169551:   "TL_CLASS_updatePtsChanged",
	1081547008:  "TL_CLASS_updateChannelWebPage",
	-686710068:  "TL_CLASS_updateDialogPinned",
	-657787251:  "TL_CLASS_updatePinnedDialogs",
	-2095595325: "TL_CLASS_updateBotWebhookJSON",
	-1684914010: "TL_CLASS_updateBotWebhookJSONQuery",
	-523384512:  "TL_CLASS_updateBotShippingQuery",
	1563376297:  "TL_CLASS_updateBotPrecheckoutQuery",
	-1425052898: "TL_CLASS_updatePhoneCall",
	281165899:   "TL_CLASS_updateLangPackTooLong",
	1442983757:  "TL_CLASS_updateLangPack",
	-451831443:  "TL_CLASS_updateFavedStickers",
	-1987495099: "TL_CLASS_updateChannelReadMessagesContents",
	1887741886:  "TL_CLASS_updateContactsReset",
	1893427255:  "TL_CLASS_updateChannelAvailableMessages",
	-1519637954: "TL_CLASS_updates_state",
	1567990072:  "TL_CLASS_updates_differenceEmpty",
	16030880:    "TL_CLASS_updates_difference",
	-1459938943: "TL_CLASS_updates_differenceSlice",
	1258196845:  "TL_CLASS_updates_differenceTooLong",
	-484987010:  "TL_CLASS_updatesTooLong",
	-1857044719: "TL_CLASS_updateShortMessage",
	377562760:   "TL_CLASS_updateShortChatMessage",
	2027216577:  "TL_CLASS_updateShort",
	1918567619:  "TL_CLASS_updatesCombined",
	1957577280:  "TL_CLASS_updates",
	301019932:   "TL_CLASS_updateShortSentMessage",
	-1916114267: "TL_CLASS_photos_photos",
	352657236:   "TL_CLASS_photos_photosSlice",
	539045032:   "TL_CLASS_photos_photo",
	157948117:   "TL_CLASS_upload_file",
	-363659686:  "TL_CLASS_upload_fileCdnRedirect",
	98092748:    "TL_CLASS_dcOption",
	-1669068444: "TL_CLASS_config",
	-1910892683: "TL_CLASS_nearestDc",
	-1987579119: "TL_CLASS_help_appUpdate",
	-1000708810: "TL_CLASS_help_noAppUpdate",
	415997816:   "TL_CLASS_help_inviteText",
	-1417756512: "TL_CLASS_encryptedChatEmpty",
	1006044124:  "TL_CLASS_encryptedChatWaiting",
	-931638658:  "TL_CLASS_encryptedChatRequested",
	-94974410:   "TL_CLASS_encryptedChat",
	332848423:   "TL_CLASS_encryptedChatDiscarded",
	-247351839:  "TL_CLASS_inputEncryptedChat",
	-1038136962: "TL_CLASS_encryptedFileEmpty",
	1248893260:  "TL_CLASS_encryptedFile",
	406307684:   "TL_CLASS_inputEncryptedFileEmpty",
	1690108678:  "TL_CLASS_inputEncryptedFileUploaded",
	1511503333:  "TL_CLASS_inputEncryptedFile",
	767652808:   "TL_CLASS_inputEncryptedFileBigUploaded",
	-317144808:  "TL_CLASS_encryptedMessage",
	594758406:   "TL_CLASS_encryptedMessageService",
	-1058912715: "TL_CLASS_messages_dhConfigNotModified",
	740433629:   "TL_CLASS_messages_dhConfig",
	1443858741:  "TL_CLASS_messages_sentEncryptedMessage",
	-1802240206: "TL_CLASS_messages_sentEncryptedFile",
	1928391342:  "TL_CLASS_inputDocumentEmpty",
	410618194:   "TL_CLASS_inputDocument",
	922273905:   "TL_CLASS_documentEmpty",
	-2027738169: "TL_CLASS_document",
	398898678:   "TL_CLASS_help_support",
	-1613493288: "TL_CLASS_notifyPeer",
	-1261946036: "TL_CLASS_notifyUsers",
	-1073230141: "TL_CLASS_notifyChats",
	1959820384:  "TL_CLASS_notifyAll",
	381645902:   "TL_CLASS_sendMessageTypingAction",
	-44119819:   "TL_CLASS_sendMessageCancelAction",
	-1584933265: "TL_CLASS_sendMessageRecordVideoAction",
	-378127636:  "TL_CLASS_sendMessageUploadVideoAction",
	-718310409:  "TL_CLASS_sendMessageRecordAudioAction",
	-212740181:  "TL_CLASS_sendMessageUploadAudioAction",
	-774682074:  "TL_CLASS_sendMessageUploadPhotoAction",
	-1441998364: "TL_CLASS_sendMessageUploadDocumentAction",
	393186209:   "TL_CLASS_sendMessageGeoLocationAction",
	1653390447:  "TL_CLASS_sendMessageChooseContactAction",
	-580219064:  "TL_CLASS_sendMessageGamePlayAction",
	-1997373508: "TL_CLASS_sendMessageRecordRoundAction",
	608050278:   "TL_CLASS_sendMessageUploadRoundAction",
	446822276:   "TL_CLASS_contacts_found",
	1335282456:  "TL_CLASS_inputPrivacyKeyStatusTimestamp",
	-1107622874: "TL_CLASS_inputPrivacyKeyChatInvite",
	-88417185:   "TL_CLASS_inputPrivacyKeyPhoneCall",
	-1137792208: "TL_CLASS_privacyKeyStatusTimestamp",
	1343122938:  "TL_CLASS_privacyKeyChatInvite",
	1030105979:  "TL_CLASS_privacyKeyPhoneCall",
	218751099:   "TL_CLASS_inputPrivacyValueAllowContacts",
	407582158:   "TL_CLASS_inputPrivacyValueAllowAll",
	320652927:   "TL_CLASS_inputPrivacyValueAllowUsers",
	195371015:   "TL_CLASS_inputPrivacyValueDisallowContacts",
	-697604407:  "TL_CLASS_inputPrivacyValueDisallowAll",
	-1877932953: "TL_CLASS_inputPrivacyValueDisallowUsers",
	-123988:     "TL_CLASS_privacyValueAllowContacts",
	1698855810:  "TL_CLASS_privacyValueAllowAll",
	1297858060:  "TL_CLASS_privacyValueAllowUsers",
	-125240806:  "TL_CLASS_privacyValueDisallowContacts",
	-1955338397: "TL_CLASS_privacyValueDisallowAll",
	209668535:   "TL_CLASS_privacyValueDisallowUsers",
	1430961007:  "TL_CLASS_account_privacyRules",
	-1194283041: "TL_CLASS_accountDaysTTL",
	1815593308:  "TL_CLASS_documentAttributeImageSize",
	297109817:   "TL_CLASS_documentAttributeAnimated",
	1662637586:  "TL_CLASS_documentAttributeSticker",
	250621158:   "TL_CLASS_documentAttributeVideo",
	-1739392570: "TL_CLASS_documentAttributeAudio",
	358154344:   "TL_CLASS_documentAttributeFilename",
	-1744710921: "TL_CLASS_documentAttributeHasStickers",
	-244016606:  "TL_CLASS_messages_stickersNotModified",
	-1970352846: "TL_CLASS_messages_stickers",
	313694676:   "TL_CLASS_stickerPack",
	-395967805:  "TL_CLASS_messages_allStickersNotModified",
	-302170017:  "TL_CLASS_messages_allStickers",
	-1369215196: "TL_CLASS_disabledFeature",
	-2066640507: "TL_CLASS_messages_affectedMessages",
	1599050311:  "TL_CLASS_contactLinkUnknown",
	-17968211:   "TL_CLASS_contactLinkNone",
	646922073:   "TL_CLASS_contactLinkHasPhone",
	-721239344:  "TL_CLASS_contactLinkContact",
	-350980120:  "TL_CLASS_webPageEmpty",
	-981018084:  "TL_CLASS_webPagePending",
	1594340540:  "TL_CLASS_webPage",
	-2054908813: "TL_CLASS_webPageNotModified",
	2079516406:  "TL_CLASS_authorization",
	307276766:   "TL_CLASS_account_authorizations",
	-1764049896: "TL_CLASS_account_noPassword",
	2081952796:  "TL_CLASS_account_password",
	-1212732749: "TL_CLASS_account_passwordSettings",
	-2037289493: "TL_CLASS_account_passwordInputSettings",
	326715557:   "TL_CLASS_auth_passwordRecovery",
	-1551583367: "TL_CLASS_receivedNotifyMessage",
	1776236393:  "TL_CLASS_chatInviteEmpty",
	-64092740:   "TL_CLASS_chatInviteExported",
	1516793212:  "TL_CLASS_chatInviteAlready",
	-613092008:  "TL_CLASS_chatInvite",
	-4838507:    "TL_CLASS_inputStickerSetEmpty",
	-1645763991: "TL_CLASS_inputStickerSetID",
	-2044933984: "TL_CLASS_inputStickerSetShortName",
	-852477119:  "TL_CLASS_stickerSet",
	-1240849242: "TL_CLASS_messages_stickerSet",
	-1032140601: "TL_CLASS_botCommand",
	-1729618630: "TL_CLASS_botInfo",
	-1560655744: "TL_CLASS_keyboardButton",
	629866245:   "TL_CLASS_keyboardButtonUrl",
	1748655686:  "TL_CLASS_keyboardButtonCallback",
	-1318425559: "TL_CLASS_keyboardButtonRequestPhone",
	-59151553:   "TL_CLASS_keyboardButtonRequestGeoLocation",
	90744648:    "TL_CLASS_keyboardButtonSwitchInline",
	1358175439:  "TL_CLASS_keyboardButtonGame",
	-1344716869: "TL_CLASS_keyboardButtonBuy",
	2002815875:  "TL_CLASS_keyboardButtonRow",
	-1606526075: "TL_CLASS_replyKeyboardHide",
	-200242528:  "TL_CLASS_replyKeyboardForceReply",
	889353612:   "TL_CLASS_replyKeyboardMarkup",
	1218642516:  "TL_CLASS_replyInlineMarkup",
	-1148011883: "TL_CLASS_messageEntityUnknown",
	-100378723:  "TL_CLASS_messageEntityMention",
	1868782349:  "TL_CLASS_messageEntityHashtag",
	1827637959:  "TL_CLASS_messageEntityBotCommand",
	1859134776:  "TL_CLASS_messageEntityUrl",
	1692693954:  "TL_CLASS_messageEntityEmail",
	-1117713463: "TL_CLASS_messageEntityBold",
	-2106619040: "TL_CLASS_messageEntityItalic",
	681706865:   "TL_CLASS_messageEntityCode",
	1938967520:  "TL_CLASS_messageEntityPre",
	1990644519:  "TL_CLASS_messageEntityTextUrl",
	892193368:   "TL_CLASS_messageEntityMentionName",
	546203849:   "TL_CLASS_inputMessageEntityMentionName",
	-292807034:  "TL_CLASS_inputChannelEmpty",
	-1343524562: "TL_CLASS_inputChannel",
	2131196633:  "TL_CLASS_contacts_resolvedPeer",
	182649427:   "TL_CLASS_messageRange",
	1041346555:  "TL_CLASS_updates_channelDifferenceEmpty",
	1788705589:  "TL_CLASS_updates_channelDifferenceTooLong",
	543450958:   "TL_CLASS_updates_channelDifference",
	-1798033689: "TL_CLASS_channelMessagesFilterEmpty",
	-847783593:  "TL_CLASS_channelMessagesFilter",
	367766557:   "TL_CLASS_channelParticipant",
	-1557620115: "TL_CLASS_channelParticipantSelf",
	-471670279:  "TL_CLASS_channelParticipantCreator",
	-1473271656: "TL_CLASS_channelParticipantAdmin",
	573315206:   "TL_CLASS_channelParticipantBanned",
	-566281095:  "TL_CLASS_channelParticipantsRecent",
	-1268741783: "TL_CLASS_channelParticipantsAdmins",
	-1548400251: "TL_CLASS_channelParticipantsKicked",
	-1328445861: "TL_CLASS_channelParticipantsBots",
	338142689:   "TL_CLASS_channelParticipantsBanned",
	106343499:   "TL_CLASS_channelParticipantsSearch",
	-177282392:  "TL_CLASS_channels_channelParticipants",
	-266911767:  "TL_CLASS_channels_channelParticipantsNotModified",
	-791039645:  "TL_CLASS_channels_channelParticipant",
	-236044656:  "TL_CLASS_help_termsOfService",
	372165663:   "TL_CLASS_foundGif",
	-1670052855: "TL_CLASS_foundGifCached",
	1158290442:  "TL_CLASS_messages_foundGifs",
	-402498398:  "TL_CLASS_messages_savedGifsNotModified",
	772213157:   "TL_CLASS_messages_savedGifs",
	691006739:   "TL_CLASS_inputBotInlineMessageMediaAuto",
	1036876423:  "TL_CLASS_inputBotInlineMessageText",
	-1045340827: "TL_CLASS_inputBotInlineMessageMediaGeo",
	-1431327288: "TL_CLASS_inputBotInlineMessageMediaVenue",
	766443943:   "TL_CLASS_inputBotInlineMessageMediaContact",
	1262639204:  "TL_CLASS_inputBotInlineMessageGame",
	750510426:   "TL_CLASS_inputBotInlineResult",
	-1462213465: "TL_CLASS_inputBotInlineResultPhoto",
	-459324:     "TL_CLASS_inputBotInlineResultDocument",
	1336154098:  "TL_CLASS_inputBotInlineResultGame",
	175419739:   "TL_CLASS_botInlineMessageMediaAuto",
	-1937807902: "TL_CLASS_botInlineMessageText",
	-1222451611: "TL_CLASS_botInlineMessageMediaGeo",
	1130767150:  "TL_CLASS_botInlineMessageMediaVenue",
	904770772:   "TL_CLASS_botInlineMessageMediaContact",
	-1679053127: "TL_CLASS_botInlineResult",
	400266251:   "TL_CLASS_botInlineMediaResult",
	-1803769784: "TL_CLASS_messages_botResults",
	524838915:   "TL_CLASS_exportedMessageLink",
	1436466797:  "TL_CLASS_messageFwdHeader",
	1923290508:  "TL_CLASS_auth_codeTypeSms",
	1948046307:  "TL_CLASS_auth_codeTypeCall",
	577556219:   "TL_CLASS_auth_codeTypeFlashCall",
	1035688326:  "TL_CLASS_auth_sentCodeTypeApp",
	-1073693790: "TL_CLASS_auth_sentCodeTypeSms",
	1398007207:  "TL_CLASS_auth_sentCodeTypeCall",
	-1425815847: "TL_CLASS_auth_sentCodeTypeFlashCall",
	911761060:   "TL_CLASS_messages_botCallbackAnswer",
	649453030:   "TL_CLASS_messages_messageEditData",
	-1995686519: "TL_CLASS_inputBotInlineMessageID",
	1008755359:  "TL_CLASS_inlineBotSwitchPM",
	863093588:   "TL_CLASS_messages_peerDialogs",
	-305282981:  "TL_CLASS_topPeer",
	-1419371685: "TL_CLASS_topPeerCategoryBotsPM",
	344356834:   "TL_CLASS_topPeerCategoryBotsInline",
	104314861:   "TL_CLASS_topPeerCategoryCorrespondents",
	-1122524854: "TL_CLASS_topPeerCategoryGroups",
	371037736:   "TL_CLASS_topPeerCategoryChannels",
	511092620:   "TL_CLASS_topPeerCategoryPhoneCalls",
	-75283823:   "TL_CLASS_topPeerCategoryPeers",
	-567906571:  "TL_CLASS_contacts_topPeersNotModified",
	1891070632:  "TL_CLASS_contacts_topPeers",
	-1169445179: "TL_CLASS_draftMessageEmpty",
	-40996577:   "TL_CLASS_draftMessage",
	82699215:    "TL_CLASS_messages_featuredStickersNotModified",
	-123893531:  "TL_CLASS_messages_featuredStickers",
	186120336:   "TL_CLASS_messages_recentStickersNotModified",
	1558317424:  "TL_CLASS_messages_recentStickers",
	1338747336:  "TL_CLASS_messages_archivedStickers",
	946083368:   "TL_CLASS_messages_stickerSetInstallResultSuccess",
	904138920:   "TL_CLASS_messages_stickerSetInstallResultArchive",
	1678812626:  "TL_CLASS_stickerSetCovered",
	872932635:   "TL_CLASS_stickerSetMultiCovered",
	-1361650766: "TL_CLASS_maskCoords",
	1251549527:  "TL_CLASS_inputStickeredMediaPhoto",
	70813275:    "TL_CLASS_inputStickeredMediaDocument",
	-1107729093: "TL_CLASS_game",
	53231223:    "TL_CLASS_inputGameID",
	-1020139510: "TL_CLASS_inputGameShortName",
	1493171408:  "TL_CLASS_highScore",
	-1707344487: "TL_CLASS_messages_highScores",
	-599948721:  "TL_CLASS_textEmpty",
	1950782688:  "TL_CLASS_textPlain",
	1730456516:  "TL_CLASS_textBold",
	-653089380:  "TL_CLASS_textItalic",
	-1054465340: "TL_CLASS_textUnderline",
	-1678197867: "TL_CLASS_textStrike",
	1816074681:  "TL_CLASS_textFixed",
	1009288385:  "TL_CLASS_textUrl",
	-564523562:  "TL_CLASS_textEmail",
	2120376535:  "TL_CLASS_textConcat",
	324435594:   "TL_CLASS_pageBlockUnsupported",
	1890305021:  "TL_CLASS_pageBlockTitle",
	-1879401953: "TL_CLASS_pageBlockSubtitle",
	-1162877472: "TL_CLASS_pageBlockAuthorDate",
	-1076861716: "TL_CLASS_pageBlockHeader",
	-248793375:  "TL_CLASS_pageBlockSubheader",
	1182402406:  "TL_CLASS_pageBlockParagraph",
	-1066346178: "TL_CLASS_pageBlockPreformatted",
	1216809369:  "TL_CLASS_pageBlockFooter",
	-618614392:  "TL_CLASS_pageBlockDivider",
	-837994576:  "TL_CLASS_pageBlockAnchor",
	978896884:   "TL_CLASS_pageBlockList",
	641563686:   "TL_CLASS_pageBlockBlockquote",
	1329878739:  "TL_CLASS_pageBlockPullquote",
	-372860542:  "TL_CLASS_pageBlockPhoto",
	-640214938:  "TL_CLASS_pageBlockVideo",
	972174080:   "TL_CLASS_pageBlockCover",
	-840826671:  "TL_CLASS_pageBlockEmbed",
	690781161:   "TL_CLASS_pageBlockEmbedPost",
	145955919:   "TL_CLASS_pageBlockCollage",
	319588707:   "TL_CLASS_pageBlockSlideshow",
	-283684427:  "TL_CLASS_pageBlockChannel",
	834148991:   "TL_CLASS_pageBlockAudio",
	-1908433218: "TL_CLASS_pagePart",
	1433323434:  "TL_CLASS_pageFull",
	-2048646399: "TL_CLASS_phoneCallDiscardReasonMissed",
	-527056480:  "TL_CLASS_phoneCallDiscardReasonDisconnect",
	1471006352:  "TL_CLASS_phoneCallDiscardReasonHangup",
	-84416311:   "TL_CLASS_phoneCallDiscardReasonBusy",
	2104790276:  "TL_CLASS_dataJSON",
	-886477832:  "TL_CLASS_labeledPrice",
	-1022713000: "TL_CLASS_invoice",
	-368917890:  "TL_CLASS_paymentCharge",
	512535275:   "TL_CLASS_postAddress",
	-1868808300: "TL_CLASS_paymentRequestedInfo",
	-842892769:  "TL_CLASS_paymentSavedCredentialsCard",
	-971322408:  "TL_CLASS_webDocument",
	-1678949555: "TL_CLASS_inputWebDocument",
	-1036396922: "TL_CLASS_inputWebFileLocation",
	568808380:   "TL_CLASS_upload_webFile",
	1062645411:  "TL_CLASS_payments_paymentForm",
	-784000893:  "TL_CLASS_payments_validatedRequestedInfo",
	1314881805:  "TL_CLASS_payments_paymentResult",
	1800845601:  "TL_CLASS_payments_paymentVerficationNeeded",
	1342771681:  "TL_CLASS_payments_paymentReceipt",
	-74456004:   "TL_CLASS_payments_savedInfo",
	-1056001329: "TL_CLASS_inputPaymentCredentialsSaved",
	873977640:   "TL_CLASS_inputPaymentCredentials",
	178373535:   "TL_CLASS_inputPaymentCredentialsApplePay",
	2035705766:  "TL_CLASS_inputPaymentCredentialsAndroidPay",
	-614138572:  "TL_CLASS_account_tmpPassword",
	-1239335713: "TL_CLASS_shippingOption",
	-6249322:    "TL_CLASS_inputStickerSetItem",
	506920429:   "TL_CLASS_inputPhoneCall",
	1399245077:  "TL_CLASS_phoneCallEmpty",
	462375633:   "TL_CLASS_phoneCallWaiting",
	-2089411356: "TL_CLASS_phoneCallRequested",
	1828732223:  "TL_CLASS_phoneCallAccepted",
	-1660057:    "TL_CLASS_phoneCall",
	1355435489:  "TL_CLASS_phoneCallDiscarded",
	-1655957568: "TL_CLASS_phoneConnection",
	-1564789301: "TL_CLASS_phoneCallProtocol",
	-326966976:  "TL_CLASS_phone_phoneCall",
	-290921362:  "TL_CLASS_upload_cdnFileReuploadNeeded",
	-1449145777: "TL_CLASS_upload_cdnFile",
	-914167110:  "TL_CLASS_cdnPublicKey",
	1462101002:  "TL_CLASS_cdnConfig",
	-892239370:  "TL_CLASS_langPackString",
	1816636575:  "TL_CLASS_langPackStringPluralized",
	695856818:   "TL_CLASS_langPackStringDeleted",
	-209337866:  "TL_CLASS_langPackDifference",
	292985073:   "TL_CLASS_langPackLanguage",
	1568467877:  "TL_CLASS_channelAdminRights",
	1489977929:  "TL_CLASS_channelBannedRights",
	-421545947:  "TL_CLASS_channelAdminLogEventActionChangeTitle",
	1427671598:  "TL_CLASS_channelAdminLogEventActionChangeAbout",
	1783299128:  "TL_CLASS_channelAdminLogEventActionChangeUsername",
	-1204857405: "TL_CLASS_channelAdminLogEventActionChangePhoto",
	460916654:   "TL_CLASS_channelAdminLogEventActionToggleInvites",
	648939889:   "TL_CLASS_channelAdminLogEventActionToggleSignatures",
	-370660328:  "TL_CLASS_channelAdminLogEventActionUpdatePinned",
	1889215493:  "TL_CLASS_channelAdminLogEventActionEditMessage",
	1121994683:  "TL_CLASS_channelAdminLogEventActionDeleteMessage",
	405815507:   "TL_CLASS_channelAdminLogEventActionParticipantJoin",
	-124291086:  "TL_CLASS_channelAdminLogEventActionParticipantLeave",
	-484690728:  "TL_CLASS_channelAdminLogEventActionParticipantInvite",
	-422036098:  "TL_CLASS_channelAdminLogEventActionParticipantToggleBan",
	-714643696:  "TL_CLASS_channelAdminLogEventActionParticipantToggleAdmin",
	-1312568665: "TL_CLASS_channelAdminLogEventActionChangeStickerSet",
	1599903217:  "TL_CLASS_channelAdminLogEventActionTogglePreHistoryHidden",
	995769920:   "TL_CLASS_channelAdminLogEvent",
	-309659827:  "TL_CLASS_channels_adminLogResults",
	-368018716:  "TL_CLASS_channelAdminLogEventsFilter",
	1558266229:  "TL_CLASS_popularContact",
	2012136335:  "TL_CLASS_cdnFileHash",
	-1634752813: "TL_CLASS_messages_favedStickersNotModified",
	-209768682:  "TL_CLASS_messages_favedStickers",
	1189204285:  "TL_CLASS_recentMeUrlUnknown",
	-1917045962: "TL_CLASS_recentMeUrlUser",
	-1608834311: "TL_CLASS_recentMeUrlChat",
	-347535331:  "TL_CLASS_recentMeUrlChatInvite",
	-1140172836: "TL_CLASS_recentMeUrlStickerSet",
	235081943:   "TL_CLASS_help_recentMeUrls",
	1588230153:  "TL_CLASS_inputSingleMedia",
	-878758099:  "TL_CLASS_invokeAfterMsg",
	1036301552:  "TL_CLASS_invokeAfterMsgs",
	-951575130:  "TL_CLASS_initConnection",
	-627372787:  "TL_CLASS_invokeWithLayer",
	-1080796745: "TL_CLASS_invokeWithoutUpdates",
	1877286395:  "TL_CLASS_auth_checkPhone",
	-2035355412: "TL_CLASS_auth_sendCode",
	453408308:   "TL_CLASS_auth_signUp",
	-1126886015: "TL_CLASS_auth_signIn",
	1461180992:  "TL_CLASS_auth_logOut",
	-1616179942: "TL_CLASS_auth_resetAuthorizations",
	1998331287:  "TL_CLASS_auth_sendInvites",
	-440401971:  "TL_CLASS_auth_exportAuthorization",
	-470837741:  "TL_CLASS_auth_importAuthorization",
	-841733627:  "TL_CLASS_auth_bindTempAuthKey",
	1738800940:  "TL_CLASS_auth_importBotAuthorization",
	174260510:   "TL_CLASS_auth_checkPassword",
	-661144474:  "TL_CLASS_auth_requestPasswordRecovery",
	1319464594:  "TL_CLASS_auth_recoverPassword",
	1056025023:  "TL_CLASS_auth_resendCode",
	520357240:   "TL_CLASS_auth_cancelCode",
	-1907842680: "TL_CLASS_auth_dropTempAuthKeys",
	1669245048:  "TL_CLASS_account_registerDevice",
	1707432768:  "TL_CLASS_account_unregisterDevice",
	-2067899501: "TL_CLASS_account_updateNotifySettings",
	313765169:   "TL_CLASS_account_getNotifySettings",
	-612493497:  "TL_CLASS_account_resetNotifySettings",
	2018596725:  "TL_CLASS_account_updateProfile",
	1713919532:  "TL_CLASS_account_updateStatus",
	-1068696894: "TL_CLASS_account_getWallPapers",
	-1374118561: "TL_CLASS_account_reportPeer",
	655677548:   "TL_CLASS_account_checkUsername",
	1040964988:  "TL_CLASS_account_updateUsername",
	-623130288:  "TL_CLASS_account_getPrivacy",
	-906486552:  "TL_CLASS_account_setPrivacy",
	1099779595:  "TL_CLASS_account_deleteAccount",
	150761757:   "TL_CLASS_account_getAccountTTL",
	608323678:   "TL_CLASS_account_setAccountTTL",
	149257707:   "TL_CLASS_account_sendChangePhoneCode",
	1891839707:  "TL_CLASS_account_changePhone",
	954152242:   "TL_CLASS_account_updateDeviceLocked",
	-484392616:  "TL_CLASS_account_getAuthorizations",
	-545786948:  "TL_CLASS_account_resetAuthorization",
	1418342645:  "TL_CLASS_account_getPassword",
	-1131605573: "TL_CLASS_account_getPasswordSettings",
	-92517498:   "TL_CLASS_account_updatePasswordSettings",
	353818557:   "TL_CLASS_account_sendConfirmPhoneCode",
	1596029123:  "TL_CLASS_account_confirmPhone",
	1250046590:  "TL_CLASS_account_getTmpPassword",
	227648840:   "TL_CLASS_users_getUsers",
	-902781519:  "TL_CLASS_users_getFullUser",
	-995929106:  "TL_CLASS_contacts_getStatuses",
	-1071414113: "TL_CLASS_contacts_getContacts",
	746589157:   "TL_CLASS_contacts_importContacts",
	-1902823612: "TL_CLASS_contacts_deleteContact",
	1504393374:  "TL_CLASS_contacts_deleteContacts",
	858475004:   "TL_CLASS_contacts_block",
	-448724803:  "TL_CLASS_contacts_unblock",
	-176409329:  "TL_CLASS_contacts_getBlocked",
	-2065352905: "TL_CLASS_contacts_exportCard",
	1340184318:  "TL_CLASS_contacts_importCard",
	301470424:   "TL_CLASS_contacts_search",
	-113456221:  "TL_CLASS_contacts_resolveUsername",
	-728224331:  "TL_CLASS_contacts_getTopPeers",
	451113900:   "TL_CLASS_contacts_resetTopPeerRating",
	-2020263951: "TL_CLASS_contacts_resetSaved",
	1109588596:  "TL_CLASS_messages_getMessages",
	421243333:   "TL_CLASS_messages_getDialogs",
	-591691168:  "TL_CLASS_messages_getHistory",
	60726944:    "TL_CLASS_messages_search",
	238054714:   "TL_CLASS_messages_readHistory",
	469850889:   "TL_CLASS_messages_deleteHistory",
	-443640366:  "TL_CLASS_messages_deleteMessages",
	94983360:    "TL_CLASS_messages_receivedMessages",
	-1551737264: "TL_CLASS_messages_setTyping",
	-91733382:   "TL_CLASS_messages_sendMessage",
	-923703407:  "TL_CLASS_messages_sendMedia",
	1888354709:  "TL_CLASS_messages_forwardMessages",
	-820669733:  "TL_CLASS_messages_reportSpam",
	-1460572005: "TL_CLASS_messages_hideReportSpam",
	913498268:   "TL_CLASS_messages_getPeerSettings",
	1013621127:  "TL_CLASS_messages_getChats",
	998448230:   "TL_CLASS_messages_getFullChat",
	-599447467:  "TL_CLASS_messages_editChatTitle",
	-900957736:  "TL_CLASS_messages_editChatPhoto",
	-106911223:  "TL_CLASS_messages_addChatUser",
	-530505962:  "TL_CLASS_messages_deleteChatUser",
	164303470:   "TL_CLASS_messages_createChat",
	865483769:   "TL_CLASS_messages_forwardMessage",
	651135312:   "TL_CLASS_messages_getDhConfig",
	-162681021:  "TL_CLASS_messages_requestEncryption",
	1035731989:  "TL_CLASS_messages_acceptEncryption",
	-304536635:  "TL_CLASS_messages_discardEncryption",
	2031374829:  "TL_CLASS_messages_setEncryptedTyping",
	2135648522:  "TL_CLASS_messages_readEncryptedHistory",
	-1451792525: "TL_CLASS_messages_sendEncrypted",
	-1701831834: "TL_CLASS_messages_sendEncryptedFile",
	852769188:   "TL_CLASS_messages_sendEncryptedService",
	1436924774:  "TL_CLASS_messages_receivedQueue",
	1259113487:  "TL_CLASS_messages_reportEncryptedSpam",
	916930423:   "TL_CLASS_messages_readMessageContents",
	479598769:   "TL_CLASS_messages_getAllStickers",
	623001124:   "TL_CLASS_messages_getWebPagePreview",
	2106086025:  "TL_CLASS_messages_exportChatInvite",
	1051570619:  "TL_CLASS_messages_checkChatInvite",
	1817183516:  "TL_CLASS_messages_importChatInvite",
	639215886:   "TL_CLASS_messages_getStickerSet",
	-946871200:  "TL_CLASS_messages_installStickerSet",
	-110209570:  "TL_CLASS_messages_uninstallStickerSet",
	-421563528:  "TL_CLASS_messages_startBot",
	-993483427:  "TL_CLASS_messages_getMessagesViews",
	-326379039:  "TL_CLASS_messages_toggleChatAdmins",
	-1444503762: "TL_CLASS_messages_editChatAdmin",
	363051235:   "TL_CLASS_messages_migrateChat",
	-1640190800: "TL_CLASS_messages_searchGlobal",
	2016638777:  "TL_CLASS_messages_reorderStickerSets",
	864953444:   "TL_CLASS_messages_getDocumentByHash",
	-1080395925: "TL_CLASS_messages_searchGifs",
	-2084618926: "TL_CLASS_messages_getSavedGifs",
	846868683:   "TL_CLASS_messages_saveGif",
	1364105629:  "TL_CLASS_messages_getInlineBotResults",
	-346119674:  "TL_CLASS_messages_setInlineBotResults",
	-1318189314: "TL_CLASS_messages_sendInlineBotResult",
	-39416522:   "TL_CLASS_messages_getMessageEditData",
	97630429:    "TL_CLASS_messages_editMessage",
	-1327463869: "TL_CLASS_messages_editInlineBotMessage",
	-2130010132: "TL_CLASS_messages_getBotCallbackAnswer",
	-712043766:  "TL_CLASS_messages_setBotCallbackAnswer",
	764901049:   "TL_CLASS_messages_getPeerDialogs",
	-1137057461: "TL_CLASS_messages_saveDraft",
	1782549861:  "TL_CLASS_messages_getAllDrafts",
	766298703:   "TL_CLASS_messages_getFeaturedStickers",
	1527873830:  "TL_CLASS_messages_readFeaturedStickers",
	1587647177:  "TL_CLASS_messages_getRecentStickers",
	958863608:   "TL_CLASS_messages_saveRecentSticker",
	-1986437075: "TL_CLASS_messages_clearRecentStickers",
	1475442322:  "TL_CLASS_messages_getArchivedStickers",
	1706608543:  "TL_CLASS_messages_getMaskStickers",
	-866424884:  "TL_CLASS_messages_getAttachedStickers",
	-1896289088: "TL_CLASS_messages_setGameScore",
	363700068:   "TL_CLASS_messages_setInlineGameScore",
	-400399203:  "TL_CLASS_messages_getGameHighScores",
	258170395:   "TL_CLASS_messages_getInlineGameHighScores",
	218777796:   "TL_CLASS_messages_getCommonChats",
	-341307408:  "TL_CLASS_messages_getAllChats",
	852135825:   "TL_CLASS_messages_getWebPage",
	847887978:   "TL_CLASS_messages_toggleDialogPin",
	-1784678844: "TL_CLASS_messages_reorderPinnedDialogs",
	-497756594:  "TL_CLASS_messages_getPinnedDialogs",
	-436833542:  "TL_CLASS_messages_setBotShippingResults",
	163765653:   "TL_CLASS_messages_setBotPrecheckoutResults",
	1369162417:  "TL_CLASS_messages_uploadMedia",
	-914493408:  "TL_CLASS_messages_sendScreenshotNotification",
	567151374:   "TL_CLASS_messages_getFavedStickers",
	-1174420133: "TL_CLASS_messages_faveSticker",
	1180140658:  "TL_CLASS_messages_getUnreadMentions",
	251759059:   "TL_CLASS_messages_readMentions",
	613691874:   "TL_CLASS_messages_getRecentLocations",
	546656559:   "TL_CLASS_messages_sendMultiMedia",
	1347929239:  "TL_CLASS_messages_uploadEncryptedFile",
	-304838614:  "TL_CLASS_updates_getState",
	630429265:   "TL_CLASS_updates_getDifference",
	51854712:    "TL_CLASS_updates_getChannelDifference",
	-256159406:  "TL_CLASS_photos_updateProfilePhoto",
	1328726168:  "TL_CLASS_photos_uploadProfilePhoto",
	-2016444625: "TL_CLASS_photos_deletePhotos",
	-1848823128: "TL_CLASS_photos_getUserPhotos",
	-1291540959: "TL_CLASS_upload_saveFilePart",
	-475607115:  "TL_CLASS_upload_getFile",
	-562337987:  "TL_CLASS_upload_saveBigFilePart",
	619086221:   "TL_CLASS_upload_getWebFile",
	536919235:   "TL_CLASS_upload_getCdnFile",
	452533257:   "TL_CLASS_upload_reuploadCdnFile",
	-149567365:  "TL_CLASS_upload_getCdnFileHashes",
	-990308245:  "TL_CLASS_help_getConfig",
	531836966:   "TL_CLASS_help_getNearestDc",
	-1372724842: "TL_CLASS_help_getAppUpdate",
	1862465352:  "TL_CLASS_help_saveAppLog",
	1295590211:  "TL_CLASS_help_getInviteText",
	-1663104819: "TL_CLASS_help_getSupport",
	-1877938321: "TL_CLASS_help_getAppChangelog",
	889286899:   "TL_CLASS_help_getTermsOfService",
	-333262899:  "TL_CLASS_help_setBotUpdatesStatus",
	1375900482:  "TL_CLASS_help_getCdnConfig",
	1036054804:  "TL_CLASS_help_getRecentMeUrls",
	-871347913:  "TL_CLASS_channels_readHistory",
	-2067661490: "TL_CLASS_channels_deleteMessages",
	-787622117:  "TL_CLASS_channels_deleteUserHistory",
	-32999408:   "TL_CLASS_channels_reportSpam",
	-1814580409: "TL_CLASS_channels_getMessages",
	306054633:   "TL_CLASS_channels_getParticipants",
	1416484774:  "TL_CLASS_channels_getParticipant",
	176122811:   "TL_CLASS_channels_getChannels",
	141781513:   "TL_CLASS_channels_getFullChannel",
	-192332417:  "TL_CLASS_channels_createChannel",
	333610782:   "TL_CLASS_channels_editAbout",
	548962836:   "TL_CLASS_channels_editAdmin",
	1450044624:  "TL_CLASS_channels_editTitle",
	-248621111:  "TL_CLASS_channels_editPhoto",
	283557164:   "TL_CLASS_channels_checkUsername",
	890549214:   "TL_CLASS_channels_updateUsername",
	615851205:   "TL_CLASS_channels_joinChannel",
	-130635115:  "TL_CLASS_channels_leaveChannel",
	429865580:   "TL_CLASS_channels_inviteToChannel",
	-950663035:  "TL_CLASS_channels_exportInvite",
	-1072619549: "TL_CLASS_channels_deleteChannel",
	1231065863:  "TL_CLASS_channels_toggleInvites",
	-934882771:  "TL_CLASS_channels_exportMessageLink",
	527021574:   "TL_CLASS_channels_toggleSignatures",
	-1490162350: "TL_CLASS_channels_updatePinnedMessage",
	-1920105769: "TL_CLASS_channels_getAdminedPublicChannels",
	-1076292147: "TL_CLASS_channels_editBanned",
	870184064:   "TL_CLASS_channels_getAdminLog",
	-359881479:  "TL_CLASS_channels_setStickers",
	-357180360:  "TL_CLASS_channels_readMessageContents",
	-1355375294: "TL_CLASS_channels_deleteHistory",
	-356796084:  "TL_CLASS_channels_togglePreHistoryHidden",
	-1440257555: "TL_CLASS_bots_sendCustomRequest",
	-434028723:  "TL_CLASS_bots_answerWebhookJSONQuery",
	-1712285883: "TL_CLASS_payments_getPaymentForm",
	-1601001088: "TL_CLASS_payments_getPaymentReceipt",
	1997180532:  "TL_CLASS_payments_validateRequestedInfo",
	730364339:   "TL_CLASS_payments_sendPaymentForm",
	578650699:   "TL_CLASS_payments_getSavedInfo",
	-667062079:  "TL_CLASS_payments_clearSavedInfo",
	-1680314774: "TL_CLASS_stickers_createStickerSet",
	-143257775:  "TL_CLASS_stickers_removeStickerFromSet",
	-4795190:    "TL_CLASS_stickers_changeStickerPosition",
	-2041315650: "TL_CLASS_stickers_addStickerToSet",
	1430593449:  "TL_CLASS_phone_getCallConfig",
	1536537556:  "TL_CLASS_phone_requestCall",
	1003664544:  "TL_CLASS_phone_acceptCall",
	788404002:   "TL_CLASS_phone_confirmCall",
	399855457:   "TL_CLASS_phone_receivedCall",
	2027164582:  "TL_CLASS_phone_discardCall",
	475228724:   "TL_CLASS_phone_setCallRating",
	662363518:   "TL_CLASS_phone_saveCallDebug",
	-1699363442: "TL_CLASS_langpack_getLangPack",
	773776152:   "TL_CLASS_langpack_getStrings",
	187583869:   "TL_CLASS_langpack_getDifference",
	-2146445955: "TL_CLASS_langpack_getLanguages",
}

var TL_CLASS_NAME_ID = map[string]int32{
	"TL_CLASS_resPQ":                                            85337187,
	"TL_CLASS_p_q_inner_data":                                   -2083955988,
	"TL_CLASS_server_DH_params_fail":                            2043348061,
	"TL_CLASS_server_DH_params_ok":                              -790100132,
	"TL_CLASS_server_DH_inner_data":                             -1249309254,
	"TL_CLASS_client_DH_inner_data":                             1715713620,
	"TL_CLASS_dh_gen_ok":                                        1003222836,
	"TL_CLASS_dh_gen_retry":                                     1188831161,
	"TL_CLASS_dh_gen_fail":                                      -1499615742,
	"TL_CLASS_destroy_auth_key_ok":                              -161422892,
	"TL_CLASS_destroy_auth_key_none":                            178201177,
	"TL_CLASS_destroy_auth_key_fail":                            -368010477,
	"TL_CLASS_req_pq":                                           1615239032,
	"TL_CLASS_req_DH_params":                                    -686627650,
	"TL_CLASS_set_client_DH_params":                             -184262881,
	"TL_CLASS_destroy_auth_key":                                 -784117408,
	"TL_CLASS_msgs_ack":                                         1658238041,
	"TL_CLASS_bad_msg_notification":                             -1477445615,
	"TL_CLASS_bad_server_salt":                                  -307542917,
	"TL_CLASS_msgs_state_req":                                   -630588590,
	"TL_CLASS_msgs_state_info":                                  81704317,
	"TL_CLASS_msgs_all_info":                                    -1933520591,
	"TL_CLASS_msg_detailed_info":                                661470918,
	"TL_CLASS_msg_new_detailed_info":                            -2137147681,
	"TL_CLASS_msg_resend_req":                                   2105940488,
	"TL_CLASS_rpc_error":                                        558156313,
	"TL_CLASS_rpc_answer_unknown":                               1579864942,
	"TL_CLASS_rpc_answer_dropped_running":                       -847714938,
	"TL_CLASS_rpc_answer_dropped":                               -1539647305,
	"TL_CLASS_future_salt":                                      155834844,
	"TL_CLASS_future_salts":                                     -1370486635,
	"TL_CLASS_pong":                                             880243653,
	"TL_CLASS_destroy_session_ok":                               -501201412,
	"TL_CLASS_destroy_session_none":                             1658015945,
	"TL_CLASS_new_session_created":                              -1631450872,
	"TL_CLASS_http_wait":                                        -1835453025,
	"TL_CLASS_ipPort":                                           -734810765,
	"TL_CLASS_help_configSimple":                                -644365371,
	"TL_CLASS_rpc_drop_answer":                                  1491380032,
	"TL_CLASS_get_future_salts":                                 -1188971260,
	"TL_CLASS_ping":                                             2059302892,
	"TL_CLASS_ping_delay_disconnect":                            -213746804,
	"TL_CLASS_destroy_session":                                  -414113498,
	"TL_CLASS_contest_saveDeveloperInfo":                        -1705021803,
	"TL_CLASS_boolFalse":                                        -1132882121,
	"TL_CLASS_boolTrue":                                         -1720552011,
	"TL_CLASS_true":                                             1072550713,
	"TL_CLASS_vector":                                           481674261,
	"TL_CLASS_error":                                            -994444869,
	"TL_CLASS_null":                                             1450380236,
	"TL_CLASS_inputPeerEmpty":                                   2134579434,
	"TL_CLASS_inputPeerSelf":                                    2107670217,
	"TL_CLASS_inputPeerChat":                                    396093539,
	"TL_CLASS_inputPeerUser":                                    2072935910,
	"TL_CLASS_inputPeerChannel":                                 548253432,
	"TL_CLASS_inputUserEmpty":                                   -1182234929,
	"TL_CLASS_inputUserSelf":                                    -138301121,
	"TL_CLASS_inputUser":                                        -668391402,
	"TL_CLASS_inputPhoneContact":                                -208488460,
	"TL_CLASS_inputFile":                                        -181407105,
	"TL_CLASS_inputFileBig":                                     -95482955,
	"TL_CLASS_inputMediaEmpty":                                  -1771768449,
	"TL_CLASS_inputMediaUploadedPhoto":                          792191537,
	"TL_CLASS_inputMediaPhoto":                                  -2114308294,
	"TL_CLASS_inputMediaGeoPoint":                               -104578748,
	"TL_CLASS_inputMediaContact":                                -1494984313,
	"TL_CLASS_inputMediaUploadedDocument":                       -476700163,
	"TL_CLASS_inputMediaDocument":                               1523279502,
	"TL_CLASS_inputMediaVenue":                                  -1052959727,
	"TL_CLASS_inputMediaGifExternal":                            1212395773,
	"TL_CLASS_inputMediaPhotoExternal":                          153267905,
	"TL_CLASS_inputMediaDocumentExternal":                       -1225309387,
	"TL_CLASS_inputMediaGame":                                   -750828557,
	"TL_CLASS_inputMediaInvoice":                                -186607933,
	"TL_CLASS_inputMediaGeoLive":                                2065305999,
	"TL_CLASS_inputChatPhotoEmpty":                              480546647,
	"TL_CLASS_inputChatUploadedPhoto":                           -1837345356,
	"TL_CLASS_inputChatPhoto":                                   -1991004873,
	"TL_CLASS_inputGeoPointEmpty":                               -457104426,
	"TL_CLASS_inputGeoPoint":                                    -206066487,
	"TL_CLASS_inputPhotoEmpty":                                  483901197,
	"TL_CLASS_inputPhoto":                                       -74070332,
	"TL_CLASS_inputFileLocation":                                342061462,
	"TL_CLASS_inputEncryptedFileLocation":                       -182231723,
	"TL_CLASS_inputDocumentFileLocation":                        1125058340,
	"TL_CLASS_inputAppEvent":                                    1996904104,
	"TL_CLASS_peerUser":                                         -1649296275,
	"TL_CLASS_peerChat":                                         -1160714821,
	"TL_CLASS_peerChannel":                                      -1109531342,
	"TL_CLASS_storage_fileUnknown":                              -1432995067,
	"TL_CLASS_storage_filePartial":                              1086091090,
	"TL_CLASS_storage_fileJpeg":                                 8322574,
	"TL_CLASS_storage_fileGif":                                  -891180321,
	"TL_CLASS_storage_filePng":                                  172975040,
	"TL_CLASS_storage_filePdf":                                  -1373745011,
	"TL_CLASS_storage_fileMp3":                                  1384777335,
	"TL_CLASS_storage_fileMov":                                  1258941372,
	"TL_CLASS_storage_fileMp4":                                  -1278304028,
	"TL_CLASS_storage_fileWebp":                                 276907596,
	"TL_CLASS_fileLocationUnavailable":                          2086234950,
	"TL_CLASS_fileLocation":                                     1406570614,
	"TL_CLASS_userEmpty":                                        537022650,
	"TL_CLASS_user":                                             773059779,
	"TL_CLASS_userProfilePhotoEmpty":                            1326562017,
	"TL_CLASS_userProfilePhoto":                                 -715532088,
	"TL_CLASS_userStatusEmpty":                                  164646985,
	"TL_CLASS_userStatusOnline":                                 -306628279,
	"TL_CLASS_userStatusOffline":                                9203775,
	"TL_CLASS_userStatusRecently":                               -496024847,
	"TL_CLASS_userStatusLastWeek":                               129960444,
	"TL_CLASS_userStatusLastMonth":                              2011940674,
	"TL_CLASS_chatEmpty":                                        -1683826688,
	"TL_CLASS_chat":                                             -652419756,
	"TL_CLASS_chatForbidden":                                    120753115,
	"TL_CLASS_channel":                                          1158377749,
	"TL_CLASS_channelForbidden":                                 681420594,
	"TL_CLASS_chatFull":                                         771925524,
	"TL_CLASS_channelFull":                                      1991201921,
	"TL_CLASS_chatParticipant":                                  -925415106,
	"TL_CLASS_chatParticipantCreator":                           -636267638,
	"TL_CLASS_chatParticipantAdmin":                             -489233354,
	"TL_CLASS_chatParticipantsForbidden":                        -57668565,
	"TL_CLASS_chatParticipants":                                 1061556205,
	"TL_CLASS_chatPhotoEmpty":                                   935395612,
	"TL_CLASS_chatPhoto":                                        1632839530,
	"TL_CLASS_messageEmpty":                                     -2082087340,
	"TL_CLASS_message":                                          1157215293,
	"TL_CLASS_messageService":                                   -1642487306,
	"TL_CLASS_messageMediaEmpty":                                1038967584,
	"TL_CLASS_messageMediaPhoto":                                -1256047857,
	"TL_CLASS_messageMediaGeo":                                  1457575028,
	"TL_CLASS_messageMediaContact":                              1585262393,
	"TL_CLASS_messageMediaUnsupported":                          -1618676578,
	"TL_CLASS_messageMediaDocument":                             2084836563,
	"TL_CLASS_messageMediaWebPage":                              -1557277184,
	"TL_CLASS_messageMediaVenue":                                784356159,
	"TL_CLASS_messageMediaGame":                                 -38694904,
	"TL_CLASS_messageMediaInvoice":                              -2074799289,
	"TL_CLASS_messageMediaGeoLive":                              2084316681,
	"TL_CLASS_messageActionEmpty":                               -1230047312,
	"TL_CLASS_messageActionChatCreate":                          -1503425638,
	"TL_CLASS_messageActionChatEditTitle":                       -1247687078,
	"TL_CLASS_messageActionChatEditPhoto":                       2144015272,
	"TL_CLASS_messageActionChatDeletePhoto":                     -1780220945,
	"TL_CLASS_messageActionChatAddUser":                         1217033015,
	"TL_CLASS_messageActionChatDeleteUser":                      -1297179892,
	"TL_CLASS_messageActionChatJoinedByLink":                    -123931160,
	"TL_CLASS_messageActionChannelCreate":                       -1781355374,
	"TL_CLASS_messageActionChatMigrateTo":                       1371385889,
	"TL_CLASS_messageActionChannelMigrateFrom":                  -1336546578,
	"TL_CLASS_messageActionPinMessage":                          -1799538451,
	"TL_CLASS_messageActionHistoryClear":                        -1615153660,
	"TL_CLASS_messageActionGameScore":                           -1834538890,
	"TL_CLASS_messageActionPaymentSentMe":                       -1892568281,
	"TL_CLASS_messageActionPaymentSent":                         1080663248,
	"TL_CLASS_messageActionPhoneCall":                           -2132731265,
	"TL_CLASS_messageActionScreenshotTaken":                     1200788123,
	"TL_CLASS_messageActionCustomAction":                        -85549226,
	"TL_CLASS_dialog":                                           -455150117,
	"TL_CLASS_photoEmpty":                                       590459437,
	"TL_CLASS_photo":                                            -1836524247,
	"TL_CLASS_photoSizeEmpty":                                   236446268,
	"TL_CLASS_photoSize":                                        2009052699,
	"TL_CLASS_photoCachedSize":                                  -374917894,
	"TL_CLASS_geoPointEmpty":                                    286776671,
	"TL_CLASS_geoPoint":                                         541710092,
	"TL_CLASS_auth_checkedPhone":                                -2128698738,
	"TL_CLASS_auth_sentCode":                                    1577067778,
	"TL_CLASS_auth_authorization":                               -855308010,
	"TL_CLASS_auth_exportedAuthorization":                       -543777747,
	"TL_CLASS_inputNotifyPeer":                                  -1195615476,
	"TL_CLASS_inputNotifyUsers":                                 423314455,
	"TL_CLASS_inputNotifyChats":                                 1251338318,
	"TL_CLASS_inputNotifyAll":                                   -1540769658,
	"TL_CLASS_inputPeerNotifyEventsEmpty":                       -265263912,
	"TL_CLASS_inputPeerNotifyEventsAll":                         -395694988,
	"TL_CLASS_inputPeerNotifySettings":                          949182130,
	"TL_CLASS_peerNotifyEventsEmpty":                            -1378534221,
	"TL_CLASS_peerNotifyEventsAll":                              1830677896,
	"TL_CLASS_peerNotifySettingsEmpty":                          1889961234,
	"TL_CLASS_peerNotifySettings":                               -1697798976,
	"TL_CLASS_peerSettings":                                     -2122045747,
	"TL_CLASS_wallPaper":                                        -860866985,
	"TL_CLASS_wallPaperSolid":                                   1662091044,
	"TL_CLASS_inputReportReasonSpam":                            1490799288,
	"TL_CLASS_inputReportReasonViolence":                        505595789,
	"TL_CLASS_inputReportReasonPornography":                     777640226,
	"TL_CLASS_inputReportReasonOther":                           -512463606,
	"TL_CLASS_userFull":                                         253890367,
	"TL_CLASS_contact":                                          -116274796,
	"TL_CLASS_importedContact":                                  -805141448,
	"TL_CLASS_contactBlocked":                                   1444661369,
	"TL_CLASS_contactStatus":                                    -748155807,
	"TL_CLASS_contacts_link":                                    986597452,
	"TL_CLASS_contacts_contactsNotModified":                     -1219778094,
	"TL_CLASS_contacts_contacts":                                -353862078,
	"TL_CLASS_contacts_importedContacts":                        2010127419,
	"TL_CLASS_contacts_blocked":                                 471043349,
	"TL_CLASS_contacts_blockedSlice":                            -1878523231,
	"TL_CLASS_messages_dialogs":                                 364538944,
	"TL_CLASS_messages_dialogsSlice":                            1910543603,
	"TL_CLASS_messages_messages":                                -1938715001,
	"TL_CLASS_messages_messagesSlice":                           189033187,
	"TL_CLASS_messages_channelMessages":                         -1725551049,
	"TL_CLASS_messages_messagesNotModified":                     1951620897,
	"TL_CLASS_messages_chats":                                   1694474197,
	"TL_CLASS_messages_chatsSlice":                              -1663561404,
	"TL_CLASS_messages_chatFull":                                -438840932,
	"TL_CLASS_messages_affectedHistory":                         -1269012015,
	"TL_CLASS_inputMessagesFilterEmpty":                         1474492012,
	"TL_CLASS_inputMessagesFilterPhotos":                        -1777752804,
	"TL_CLASS_inputMessagesFilterVideo":                         -1614803355,
	"TL_CLASS_inputMessagesFilterPhotoVideo":                    1458172132,
	"TL_CLASS_inputMessagesFilterDocument":                      -1629621880,
	"TL_CLASS_inputMessagesFilterUrl":                           2129714567,
	"TL_CLASS_inputMessagesFilterGif":                           -3644025,
	"TL_CLASS_inputMessagesFilterVoice":                         1358283666,
	"TL_CLASS_inputMessagesFilterMusic":                         928101534,
	"TL_CLASS_inputMessagesFilterChatPhotos":                    975236280,
	"TL_CLASS_inputMessagesFilterPhoneCalls":                    -2134272152,
	"TL_CLASS_inputMessagesFilterRoundVoice":                    2054952868,
	"TL_CLASS_inputMessagesFilterRoundVideo":                    -1253451181,
	"TL_CLASS_inputMessagesFilterMyMentions":                    -1040652646,
	"TL_CLASS_inputMessagesFilterGeo":                           -419271411,
	"TL_CLASS_inputMessagesFilterContacts":                      -530392189,
	"TL_CLASS_updateNewMessage":                                 522914557,
	"TL_CLASS_updateMessageID":                                  1318109142,
	"TL_CLASS_updateDeleteMessages":                             -1576161051,
	"TL_CLASS_updateUserTyping":                                 1548249383,
	"TL_CLASS_updateChatUserTyping":                             -1704596961,
	"TL_CLASS_updateChatParticipants":                           125178264,
	"TL_CLASS_updateUserStatus":                                 469489699,
	"TL_CLASS_updateUserName":                                   -1489818765,
	"TL_CLASS_updateUserPhoto":                                  -1791935732,
	"TL_CLASS_updateContactRegistered":                          628472761,
	"TL_CLASS_updateContactLink":                                -1657903163,
	"TL_CLASS_updateNewEncryptedMessage":                        314359194,
	"TL_CLASS_updateEncryptedChatTyping":                        386986326,
	"TL_CLASS_updateEncryption":                                 -1264392051,
	"TL_CLASS_updateEncryptedMessagesRead":                      956179895,
	"TL_CLASS_updateChatParticipantAdd":                         -364179876,
	"TL_CLASS_updateChatParticipantDelete":                      1851755554,
	"TL_CLASS_updateDcOptions":                                  -1906403213,
	"TL_CLASS_updateUserBlocked":                                -2131957734,
	"TL_CLASS_updateNotifySettings":                             -1094555409,
	"TL_CLASS_updateServiceNotification":                        -337352679,
	"TL_CLASS_updatePrivacy":                                    -298113238,
	"TL_CLASS_updateUserPhone":                                  314130811,
	"TL_CLASS_updateReadHistoryInbox":                           -1721631396,
	"TL_CLASS_updateReadHistoryOutbox":                          791617983,
	"TL_CLASS_updateWebPage":                                    2139689491,
	"TL_CLASS_updateReadMessagesContents":                       1757493555,
	"TL_CLASS_updateChannelTooLong":                             -352032773,
	"TL_CLASS_updateChannel":                                    -1227598250,
	"TL_CLASS_updateNewChannelMessage":                          1656358105,
	"TL_CLASS_updateReadChannelInbox":                           1108669311,
	"TL_CLASS_updateDeleteChannelMessages":                      -1015733815,
	"TL_CLASS_updateChannelMessageViews":                        -1734268085,
	"TL_CLASS_updateChatAdmins":                                 1855224129,
	"TL_CLASS_updateChatParticipantAdmin":                       -1232070311,
	"TL_CLASS_updateNewStickerSet":                              1753886890,
	"TL_CLASS_updateStickerSetsOrder":                           196268545,
	"TL_CLASS_updateStickerSets":                                1135492588,
	"TL_CLASS_updateSavedGifs":                                  -1821035490,
	"TL_CLASS_updateBotInlineQuery":                             1417832080,
	"TL_CLASS_updateBotInlineSend":                              239663460,
	"TL_CLASS_updateEditChannelMessage":                         457133559,
	"TL_CLASS_updateChannelPinnedMessage":                       -1738988427,
	"TL_CLASS_updateBotCallbackQuery":                           -415938591,
	"TL_CLASS_updateEditMessage":                                -469536605,
	"TL_CLASS_updateInlineBotCallbackQuery":                     -103646630,
	"TL_CLASS_updateReadChannelOutbox":                          634833351,
	"TL_CLASS_updateDraftMessage":                               -299124375,
	"TL_CLASS_updateReadFeaturedStickers":                       1461528386,
	"TL_CLASS_updateRecentStickers":                             -1706939360,
	"TL_CLASS_updateConfig":                                     -1574314746,
	"TL_CLASS_updatePtsChanged":                                 861169551,
	"TL_CLASS_updateChannelWebPage":                             1081547008,
	"TL_CLASS_updateDialogPinned":                               -686710068,
	"TL_CLASS_updatePinnedDialogs":                              -657787251,
	"TL_CLASS_updateBotWebhookJSON":                             -2095595325,
	"TL_CLASS_updateBotWebhookJSONQuery":                        -1684914010,
	"TL_CLASS_updateBotShippingQuery":                           -523384512,
	"TL_CLASS_updateBotPrecheckoutQuery":                        1563376297,
	"TL_CLASS_updatePhoneCall":                                  -1425052898,
	"TL_CLASS_updateLangPackTooLong":                            281165899,
	"TL_CLASS_updateLangPack":                                   1442983757,
	"TL_CLASS_updateFavedStickers":                              -451831443,
	"TL_CLASS_updateChannelReadMessagesContents":                -1987495099,
	"TL_CLASS_updateContactsReset":                              1887741886,
	"TL_CLASS_updateChannelAvailableMessages":                   1893427255,
	"TL_CLASS_updates_state":                                    -1519637954,
	"TL_CLASS_updates_differenceEmpty":                          1567990072,
	"TL_CLASS_updates_difference":                               16030880,
	"TL_CLASS_updates_differenceSlice":                          -1459938943,
	"TL_CLASS_updates_differenceTooLong":                        1258196845,
	"TL_CLASS_updatesTooLong":                                   -484987010,
	"TL_CLASS_updateShortMessage":                               -1857044719,
	"TL_CLASS_updateShortChatMessage":                           377562760,
	"TL_CLASS_updateShort":                                      2027216577,
	"TL_CLASS_updatesCombined":                                  1918567619,
	"TL_CLASS_updates":                                          1957577280,
	"TL_CLASS_updateShortSentMessage":                           301019932,
	"TL_CLASS_photos_photos":                                    -1916114267,
	"TL_CLASS_photos_photosSlice":                               352657236,
	"TL_CLASS_photos_photo":                                     539045032,
	"TL_CLASS_upload_file":                                      157948117,
	"TL_CLASS_upload_fileCdnRedirect":                           -363659686,
	"TL_CLASS_dcOption":                                         98092748,
	"TL_CLASS_config":                                           -1669068444,
	"TL_CLASS_nearestDc":                                        -1910892683,
	"TL_CLASS_help_appUpdate":                                   -1987579119,
	"TL_CLASS_help_noAppUpdate":                                 -1000708810,
	"TL_CLASS_help_inviteText":                                  415997816,
	"TL_CLASS_encryptedChatEmpty":                               -1417756512,
	"TL_CLASS_encryptedChatWaiting":                             1006044124,
	"TL_CLASS_encryptedChatRequested":                           -931638658,
	"TL_CLASS_encryptedChat":                                    -94974410,
	"TL_CLASS_encryptedChatDiscarded":                           332848423,
	"TL_CLASS_inputEncryptedChat":                               -247351839,
	"TL_CLASS_encryptedFileEmpty":                               -1038136962,
	"TL_CLASS_encryptedFile":                                    1248893260,
	"TL_CLASS_inputEncryptedFileEmpty":                          406307684,
	"TL_CLASS_inputEncryptedFileUploaded":                       1690108678,
	"TL_CLASS_inputEncryptedFile":                               1511503333,
	"TL_CLASS_inputEncryptedFileBigUploaded":                    767652808,
	"TL_CLASS_encryptedMessage":                                 -317144808,
	"TL_CLASS_encryptedMessageService":                          594758406,
	"TL_CLASS_messages_dhConfigNotModified":                     -1058912715,
	"TL_CLASS_messages_dhConfig":                                740433629,
	"TL_CLASS_messages_sentEncryptedMessage":                    1443858741,
	"TL_CLASS_messages_sentEncryptedFile":                       -1802240206,
	"TL_CLASS_inputDocumentEmpty":                               1928391342,
	"TL_CLASS_inputDocument":                                    410618194,
	"TL_CLASS_documentEmpty":                                    922273905,
	"TL_CLASS_document":                                         -2027738169,
	"TL_CLASS_help_support":                                     398898678,
	"TL_CLASS_notifyPeer":                                       -1613493288,
	"TL_CLASS_notifyUsers":                                      -1261946036,
	"TL_CLASS_notifyChats":                                      -1073230141,
	"TL_CLASS_notifyAll":                                        1959820384,
	"TL_CLASS_sendMessageTypingAction":                          381645902,
	"TL_CLASS_sendMessageCancelAction":                          -44119819,
	"TL_CLASS_sendMessageRecordVideoAction":                     -1584933265,
	"TL_CLASS_sendMessageUploadVideoAction":                     -378127636,
	"TL_CLASS_sendMessageRecordAudioAction":                     -718310409,
	"TL_CLASS_sendMessageUploadAudioAction":                     -212740181,
	"TL_CLASS_sendMessageUploadPhotoAction":                     -774682074,
	"TL_CLASS_sendMessageUploadDocumentAction":                  -1441998364,
	"TL_CLASS_sendMessageGeoLocationAction":                     393186209,
	"TL_CLASS_sendMessageChooseContactAction":                   1653390447,
	"TL_CLASS_sendMessageGamePlayAction":                        -580219064,
	"TL_CLASS_sendMessageRecordRoundAction":                     -1997373508,
	"TL_CLASS_sendMessageUploadRoundAction":                     608050278,
	"TL_CLASS_contacts_found":                                   446822276,
	"TL_CLASS_inputPrivacyKeyStatusTimestamp":                   1335282456,
	"TL_CLASS_inputPrivacyKeyChatInvite":                        -1107622874,
	"TL_CLASS_inputPrivacyKeyPhoneCall":                         -88417185,
	"TL_CLASS_privacyKeyStatusTimestamp":                        -1137792208,
	"TL_CLASS_privacyKeyChatInvite":                             1343122938,
	"TL_CLASS_privacyKeyPhoneCall":                              1030105979,
	"TL_CLASS_inputPrivacyValueAllowContacts":                   218751099,
	"TL_CLASS_inputPrivacyValueAllowAll":                        407582158,
	"TL_CLASS_inputPrivacyValueAllowUsers":                      320652927,
	"TL_CLASS_inputPrivacyValueDisallowContacts":                195371015,
	"TL_CLASS_inputPrivacyValueDisallowAll":                     -697604407,
	"TL_CLASS_inputPrivacyValueDisallowUsers":                   -1877932953,
	"TL_CLASS_privacyValueAllowContacts":                        -123988,
	"TL_CLASS_privacyValueAllowAll":                             1698855810,
	"TL_CLASS_privacyValueAllowUsers":                           1297858060,
	"TL_CLASS_privacyValueDisallowContacts":                     -125240806,
	"TL_CLASS_privacyValueDisallowAll":                          -1955338397,
	"TL_CLASS_privacyValueDisallowUsers":                        209668535,
	"TL_CLASS_account_privacyRules":                             1430961007,
	"TL_CLASS_accountDaysTTL":                                   -1194283041,
	"TL_CLASS_documentAttributeImageSize":                       1815593308,
	"TL_CLASS_documentAttributeAnimated":                        297109817,
	"TL_CLASS_documentAttributeSticker":                         1662637586,
	"TL_CLASS_documentAttributeVideo":                           250621158,
	"TL_CLASS_documentAttributeAudio":                           -1739392570,
	"TL_CLASS_documentAttributeFilename":                        358154344,
	"TL_CLASS_documentAttributeHasStickers":                     -1744710921,
	"TL_CLASS_messages_stickersNotModified":                     -244016606,
	"TL_CLASS_messages_stickers":                                -1970352846,
	"TL_CLASS_stickerPack":                                      313694676,
	"TL_CLASS_messages_allStickersNotModified":                  -395967805,
	"TL_CLASS_messages_allStickers":                             -302170017,
	"TL_CLASS_disabledFeature":                                  -1369215196,
	"TL_CLASS_messages_affectedMessages":                        -2066640507,
	"TL_CLASS_contactLinkUnknown":                               1599050311,
	"TL_CLASS_contactLinkNone":                                  -17968211,
	"TL_CLASS_contactLinkHasPhone":                              646922073,
	"TL_CLASS_contactLinkContact":                               -721239344,
	"TL_CLASS_webPageEmpty":                                     -350980120,
	"TL_CLASS_webPagePending":                                   -981018084,
	"TL_CLASS_webPage":                                          1594340540,
	"TL_CLASS_webPageNotModified":                               -2054908813,
	"TL_CLASS_authorization":                                    2079516406,
	"TL_CLASS_account_authorizations":                           307276766,
	"TL_CLASS_account_noPassword":                               -1764049896,
	"TL_CLASS_account_password":                                 2081952796,
	"TL_CLASS_account_passwordSettings":                         -1212732749,
	"TL_CLASS_account_passwordInputSettings":                    -2037289493,
	"TL_CLASS_auth_passwordRecovery":                            326715557,
	"TL_CLASS_receivedNotifyMessage":                            -1551583367,
	"TL_CLASS_chatInviteEmpty":                                  1776236393,
	"TL_CLASS_chatInviteExported":                               -64092740,
	"TL_CLASS_chatInviteAlready":                                1516793212,
	"TL_CLASS_chatInvite":                                       -613092008,
	"TL_CLASS_inputStickerSetEmpty":                             -4838507,
	"TL_CLASS_inputStickerSetID":                                -1645763991,
	"TL_CLASS_inputStickerSetShortName":                         -2044933984,
	"TL_CLASS_stickerSet":                                       -852477119,
	"TL_CLASS_messages_stickerSet":                              -1240849242,
	"TL_CLASS_botCommand":                                       -1032140601,
	"TL_CLASS_botInfo":                                          -1729618630,
	"TL_CLASS_keyboardButton":                                   -1560655744,
	"TL_CLASS_keyboardButtonUrl":                                629866245,
	"TL_CLASS_keyboardButtonCallback":                           1748655686,
	"TL_CLASS_keyboardButtonRequestPhone":                       -1318425559,
	"TL_CLASS_keyboardButtonRequestGeoLocation":                 -59151553,
	"TL_CLASS_keyboardButtonSwitchInline":                       90744648,
	"TL_CLASS_keyboardButtonGame":                               1358175439,
	"TL_CLASS_keyboardButtonBuy":                                -1344716869,
	"TL_CLASS_keyboardButtonRow":                                2002815875,
	"TL_CLASS_replyKeyboardHide":                                -1606526075,
	"TL_CLASS_replyKeyboardForceReply":                          -200242528,
	"TL_CLASS_replyKeyboardMarkup":                              889353612,
	"TL_CLASS_replyInlineMarkup":                                1218642516,
	"TL_CLASS_messageEntityUnknown":                             -1148011883,
	"TL_CLASS_messageEntityMention":                             -100378723,
	"TL_CLASS_messageEntityHashtag":                             1868782349,
	"TL_CLASS_messageEntityBotCommand":                          1827637959,
	"TL_CLASS_messageEntityUrl":                                 1859134776,
	"TL_CLASS_messageEntityEmail":                               1692693954,
	"TL_CLASS_messageEntityBold":                                -1117713463,
	"TL_CLASS_messageEntityItalic":                              -2106619040,
	"TL_CLASS_messageEntityCode":                                681706865,
	"TL_CLASS_messageEntityPre":                                 1938967520,
	"TL_CLASS_messageEntityTextUrl":                             1990644519,
	"TL_CLASS_messageEntityMentionName":                         892193368,
	"TL_CLASS_inputMessageEntityMentionName":                    546203849,
	"TL_CLASS_inputChannelEmpty":                                -292807034,
	"TL_CLASS_inputChannel":                                     -1343524562,
	"TL_CLASS_contacts_resolvedPeer":                            2131196633,
	"TL_CLASS_messageRange":                                     182649427,
	"TL_CLASS_updates_channelDifferenceEmpty":                   1041346555,
	"TL_CLASS_updates_channelDifferenceTooLong":                 1788705589,
	"TL_CLASS_updates_channelDifference":                        543450958,
	"TL_CLASS_channelMessagesFilterEmpty":                       -1798033689,
	"TL_CLASS_channelMessagesFilter":                            -847783593,
	"TL_CLASS_channelParticipant":                               367766557,
	"TL_CLASS_channelParticipantSelf":                           -1557620115,
	"TL_CLASS_channelParticipantCreator":                        -471670279,
	"TL_CLASS_channelParticipantAdmin":                          -1473271656,
	"TL_CLASS_channelParticipantBanned":                         573315206,
	"TL_CLASS_channelParticipantsRecent":                        -566281095,
	"TL_CLASS_channelParticipantsAdmins":                        -1268741783,
	"TL_CLASS_channelParticipantsKicked":                        -1548400251,
	"TL_CLASS_channelParticipantsBots":                          -1328445861,
	"TL_CLASS_channelParticipantsBanned":                        338142689,
	"TL_CLASS_channelParticipantsSearch":                        106343499,
	"TL_CLASS_channels_channelParticipants":                     -177282392,
	"TL_CLASS_channels_channelParticipantsNotModified":          -266911767,
	"TL_CLASS_channels_channelParticipant":                      -791039645,
	"TL_CLASS_help_termsOfService":                              -236044656,
	"TL_CLASS_foundGif":                                         372165663,
	"TL_CLASS_foundGifCached":                                   -1670052855,
	"TL_CLASS_messages_foundGifs":                               1158290442,
	"TL_CLASS_messages_savedGifsNotModified":                    -402498398,
	"TL_CLASS_messages_savedGifs":                               772213157,
	"TL_CLASS_inputBotInlineMessageMediaAuto":                   691006739,
	"TL_CLASS_inputBotInlineMessageText":                        1036876423,
	"TL_CLASS_inputBotInlineMessageMediaGeo":                    -1045340827,
	"TL_CLASS_inputBotInlineMessageMediaVenue":                  -1431327288,
	"TL_CLASS_inputBotInlineMessageMediaContact":                766443943,
	"TL_CLASS_inputBotInlineMessageGame":                        1262639204,
	"TL_CLASS_inputBotInlineResult":                             750510426,
	"TL_CLASS_inputBotInlineResultPhoto":                        -1462213465,
	"TL_CLASS_inputBotInlineResultDocument":                     -459324,
	"TL_CLASS_inputBotInlineResultGame":                         1336154098,
	"TL_CLASS_botInlineMessageMediaAuto":                        175419739,
	"TL_CLASS_botInlineMessageText":                             -1937807902,
	"TL_CLASS_botInlineMessageMediaGeo":                         -1222451611,
	"TL_CLASS_botInlineMessageMediaVenue":                       1130767150,
	"TL_CLASS_botInlineMessageMediaContact":                     904770772,
	"TL_CLASS_botInlineResult":                                  -1679053127,
	"TL_CLASS_botInlineMediaResult":                             400266251,
	"TL_CLASS_messages_botResults":                              -1803769784,
	"TL_CLASS_exportedMessageLink":                              524838915,
	"TL_CLASS_messageFwdHeader":                                 1436466797,
	"TL_CLASS_auth_codeTypeSms":                                 1923290508,
	"TL_CLASS_auth_codeTypeCall":                                1948046307,
	"TL_CLASS_auth_codeTypeFlashCall":                           577556219,
	"TL_CLASS_auth_sentCodeTypeApp":                             1035688326,
	"TL_CLASS_auth_sentCodeTypeSms":                             -1073693790,
	"TL_CLASS_auth_sentCodeTypeCall":                            1398007207,
	"TL_CLASS_auth_sentCodeTypeFlashCall":                       -1425815847,
	"TL_CLASS_messages_botCallbackAnswer":                       911761060,
	"TL_CLASS_messages_messageEditData":                         649453030,
	"TL_CLASS_inputBotInlineMessageID":                          -1995686519,
	"TL_CLASS_inlineBotSwitchPM":                                1008755359,
	"TL_CLASS_messages_peerDialogs":                             863093588,
	"TL_CLASS_topPeer":                                          -305282981,
	"TL_CLASS_topPeerCategoryBotsPM":                            -1419371685,
	"TL_CLASS_topPeerCategoryBotsInline":                        344356834,
	"TL_CLASS_topPeerCategoryCorrespondents":                    104314861,
	"TL_CLASS_topPeerCategoryGroups":                            -1122524854,
	"TL_CLASS_topPeerCategoryChannels":                          371037736,
	"TL_CLASS_topPeerCategoryPhoneCalls":                        511092620,
	"TL_CLASS_topPeerCategoryPeers":                             -75283823,
	"TL_CLASS_contacts_topPeersNotModified":                     -567906571,
	"TL_CLASS_contacts_topPeers":                                1891070632,
	"TL_CLASS_draftMessageEmpty":                                -1169445179,
	"TL_CLASS_draftMessage":                                     -40996577,
	"TL_CLASS_messages_featuredStickersNotModified":             82699215,
	"TL_CLASS_messages_featuredStickers":                        -123893531,
	"TL_CLASS_messages_recentStickersNotModified":               186120336,
	"TL_CLASS_messages_recentStickers":                          1558317424,
	"TL_CLASS_messages_archivedStickers":                        1338747336,
	"TL_CLASS_messages_stickerSetInstallResultSuccess":          946083368,
	"TL_CLASS_messages_stickerSetInstallResultArchive":          904138920,
	"TL_CLASS_stickerSetCovered":                                1678812626,
	"TL_CLASS_stickerSetMultiCovered":                           872932635,
	"TL_CLASS_maskCoords":                                       -1361650766,
	"TL_CLASS_inputStickeredMediaPhoto":                         1251549527,
	"TL_CLASS_inputStickeredMediaDocument":                      70813275,
	"TL_CLASS_game":                                             -1107729093,
	"TL_CLASS_inputGameID":                                      53231223,
	"TL_CLASS_inputGameShortName":                               -1020139510,
	"TL_CLASS_highScore":                                        1493171408,
	"TL_CLASS_messages_highScores":                              -1707344487,
	"TL_CLASS_textEmpty":                                        -599948721,
	"TL_CLASS_textPlain":                                        1950782688,
	"TL_CLASS_textBold":                                         1730456516,
	"TL_CLASS_textItalic":                                       -653089380,
	"TL_CLASS_textUnderline":                                    -1054465340,
	"TL_CLASS_textStrike":                                       -1678197867,
	"TL_CLASS_textFixed":                                        1816074681,
	"TL_CLASS_textUrl":                                          1009288385,
	"TL_CLASS_textEmail":                                        -564523562,
	"TL_CLASS_textConcat":                                       2120376535,
	"TL_CLASS_pageBlockUnsupported":                             324435594,
	"TL_CLASS_pageBlockTitle":                                   1890305021,
	"TL_CLASS_pageBlockSubtitle":                                -1879401953,
	"TL_CLASS_pageBlockAuthorDate":                              -1162877472,
	"TL_CLASS_pageBlockHeader":                                  -1076861716,
	"TL_CLASS_pageBlockSubheader":                               -248793375,
	"TL_CLASS_pageBlockParagraph":                               1182402406,
	"TL_CLASS_pageBlockPreformatted":                            -1066346178,
	"TL_CLASS_pageBlockFooter":                                  1216809369,
	"TL_CLASS_pageBlockDivider":                                 -618614392,
	"TL_CLASS_pageBlockAnchor":                                  -837994576,
	"TL_CLASS_pageBlockList":                                    978896884,
	"TL_CLASS_pageBlockBlockquote":                              641563686,
	"TL_CLASS_pageBlockPullquote":                               1329878739,
	"TL_CLASS_pageBlockPhoto":                                   -372860542,
	"TL_CLASS_pageBlockVideo":                                   -640214938,
	"TL_CLASS_pageBlockCover":                                   972174080,
	"TL_CLASS_pageBlockEmbed":                                   -840826671,
	"TL_CLASS_pageBlockEmbedPost":                               690781161,
	"TL_CLASS_pageBlockCollage":                                 145955919,
	"TL_CLASS_pageBlockSlideshow":                               319588707,
	"TL_CLASS_pageBlockChannel":                                 -283684427,
	"TL_CLASS_pageBlockAudio":                                   834148991,
	"TL_CLASS_pagePart":                                         -1908433218,
	"TL_CLASS_pageFull":                                         1433323434,
	"TL_CLASS_phoneCallDiscardReasonMissed":                     -2048646399,
	"TL_CLASS_phoneCallDiscardReasonDisconnect":                 -527056480,
	"TL_CLASS_phoneCallDiscardReasonHangup":                     1471006352,
	"TL_CLASS_phoneCallDiscardReasonBusy":                       -84416311,
	"TL_CLASS_dataJSON":                                         2104790276,
	"TL_CLASS_labeledPrice":                                     -886477832,
	"TL_CLASS_invoice":                                          -1022713000,
	"TL_CLASS_paymentCharge":                                    -368917890,
	"TL_CLASS_postAddress":                                      512535275,
	"TL_CLASS_paymentRequestedInfo":                             -1868808300,
	"TL_CLASS_paymentSavedCredentialsCard":                      -842892769,
	"TL_CLASS_webDocument":                                      -971322408,
	"TL_CLASS_inputWebDocument":                                 -1678949555,
	"TL_CLASS_inputWebFileLocation":                             -1036396922,
	"TL_CLASS_upload_webFile":                                   568808380,
	"TL_CLASS_payments_paymentForm":                             1062645411,
	"TL_CLASS_payments_validatedRequestedInfo":                  -784000893,
	"TL_CLASS_payments_paymentResult":                           1314881805,
	"TL_CLASS_payments_paymentVerficationNeeded":                1800845601,
	"TL_CLASS_payments_paymentReceipt":                          1342771681,
	"TL_CLASS_payments_savedInfo":                               -74456004,
	"TL_CLASS_inputPaymentCredentialsSaved":                     -1056001329,
	"TL_CLASS_inputPaymentCredentials":                          873977640,
	"TL_CLASS_inputPaymentCredentialsApplePay":                  178373535,
	"TL_CLASS_inputPaymentCredentialsAndroidPay":                2035705766,
	"TL_CLASS_account_tmpPassword":                              -614138572,
	"TL_CLASS_shippingOption":                                   -1239335713,
	"TL_CLASS_inputStickerSetItem":                              -6249322,
	"TL_CLASS_inputPhoneCall":                                   506920429,
	"TL_CLASS_phoneCallEmpty":                                   1399245077,
	"TL_CLASS_phoneCallWaiting":                                 462375633,
	"TL_CLASS_phoneCallRequested":                               -2089411356,
	"TL_CLASS_phoneCallAccepted":                                1828732223,
	"TL_CLASS_phoneCall":                                        -1660057,
	"TL_CLASS_phoneCallDiscarded":                               1355435489,
	"TL_CLASS_phoneConnection":                                  -1655957568,
	"TL_CLASS_phoneCallProtocol":                                -1564789301,
	"TL_CLASS_phone_phoneCall":                                  -326966976,
	"TL_CLASS_upload_cdnFileReuploadNeeded":                     -290921362,
	"TL_CLASS_upload_cdnFile":                                   -1449145777,
	"TL_CLASS_cdnPublicKey":                                     -914167110,
	"TL_CLASS_cdnConfig":                                        1462101002,
	"TL_CLASS_langPackString":                                   -892239370,
	"TL_CLASS_langPackStringPluralized":                         1816636575,
	"TL_CLASS_langPackStringDeleted":                            695856818,
	"TL_CLASS_langPackDifference":                               -209337866,
	"TL_CLASS_langPackLanguage":                                 292985073,
	"TL_CLASS_channelAdminRights":                               1568467877,
	"TL_CLASS_channelBannedRights":                              1489977929,
	"TL_CLASS_channelAdminLogEventActionChangeTitle":            -421545947,
	"TL_CLASS_channelAdminLogEventActionChangeAbout":            1427671598,
	"TL_CLASS_channelAdminLogEventActionChangeUsername":         1783299128,
	"TL_CLASS_channelAdminLogEventActionChangePhoto":            -1204857405,
	"TL_CLASS_channelAdminLogEventActionToggleInvites":          460916654,
	"TL_CLASS_channelAdminLogEventActionToggleSignatures":       648939889,
	"TL_CLASS_channelAdminLogEventActionUpdatePinned":           -370660328,
	"TL_CLASS_channelAdminLogEventActionEditMessage":            1889215493,
	"TL_CLASS_channelAdminLogEventActionDeleteMessage":          1121994683,
	"TL_CLASS_channelAdminLogEventActionParticipantJoin":        405815507,
	"TL_CLASS_channelAdminLogEventActionParticipantLeave":       -124291086,
	"TL_CLASS_channelAdminLogEventActionParticipantInvite":      -484690728,
	"TL_CLASS_channelAdminLogEventActionParticipantToggleBan":   -422036098,
	"TL_CLASS_channelAdminLogEventActionParticipantToggleAdmin": -714643696,
	"TL_CLASS_channelAdminLogEventActionChangeStickerSet":       -1312568665,
	"TL_CLASS_channelAdminLogEventActionTogglePreHistoryHidden": 1599903217,
	"TL_CLASS_channelAdminLogEvent":                             995769920,
	"TL_CLASS_channels_adminLogResults":                         -309659827,
	"TL_CLASS_channelAdminLogEventsFilter":                      -368018716,
	"TL_CLASS_popularContact":                                   1558266229,
	"TL_CLASS_cdnFileHash":                                      2012136335,
	"TL_CLASS_messages_favedStickersNotModified":                -1634752813,
	"TL_CLASS_messages_favedStickers":                           -209768682,
	"TL_CLASS_recentMeUrlUnknown":                               1189204285,
	"TL_CLASS_recentMeUrlUser":                                  -1917045962,
	"TL_CLASS_recentMeUrlChat":                                  -1608834311,
	"TL_CLASS_recentMeUrlChatInvite":                            -347535331,
	"TL_CLASS_recentMeUrlStickerSet":                            -1140172836,
	"TL_CLASS_help_recentMeUrls":                                235081943,
	"TL_CLASS_inputSingleMedia":                                 1588230153,
	"TL_CLASS_invokeAfterMsg":                                   -878758099,
	"TL_CLASS_invokeAfterMsgs":                                  1036301552,
	"TL_CLASS_initConnection":                                   -951575130,
	"TL_CLASS_invokeWithLayer":                                  -627372787,
	"TL_CLASS_invokeWithoutUpdates":                             -1080796745,
	"TL_CLASS_auth_checkPhone":                                  1877286395,
	"TL_CLASS_auth_sendCode":                                    -2035355412,
	"TL_CLASS_auth_signUp":                                      453408308,
	"TL_CLASS_auth_signIn":                                      -1126886015,
	"TL_CLASS_auth_logOut":                                      1461180992,
	"TL_CLASS_auth_resetAuthorizations":                         -1616179942,
	"TL_CLASS_auth_sendInvites":                                 1998331287,
	"TL_CLASS_auth_exportAuthorization":                         -440401971,
	"TL_CLASS_auth_importAuthorization":                         -470837741,
	"TL_CLASS_auth_bindTempAuthKey":                             -841733627,
	"TL_CLASS_auth_importBotAuthorization":                      1738800940,
	"TL_CLASS_auth_checkPassword":                               174260510,
	"TL_CLASS_auth_requestPasswordRecovery":                     -661144474,
	"TL_CLASS_auth_recoverPassword":                             1319464594,
	"TL_CLASS_auth_resendCode":                                  1056025023,
	"TL_CLASS_auth_cancelCode":                                  520357240,
	"TL_CLASS_auth_dropTempAuthKeys":                            -1907842680,
	"TL_CLASS_account_registerDevice":                           1669245048,
	"TL_CLASS_account_unregisterDevice":                         1707432768,
	"TL_CLASS_account_updateNotifySettings":                     -2067899501,
	"TL_CLASS_account_getNotifySettings":                        313765169,
	"TL_CLASS_account_resetNotifySettings":                      -612493497,
	"TL_CLASS_account_updateProfile":                            2018596725,
	"TL_CLASS_account_updateStatus":                             1713919532,
	"TL_CLASS_account_getWallPapers":                            -1068696894,
	"TL_CLASS_account_reportPeer":                               -1374118561,
	"TL_CLASS_account_checkUsername":                            655677548,
	"TL_CLASS_account_updateUsername":                           1040964988,
	"TL_CLASS_account_getPrivacy":                               -623130288,
	"TL_CLASS_account_setPrivacy":                               -906486552,
	"TL_CLASS_account_deleteAccount":                            1099779595,
	"TL_CLASS_account_getAccountTTL":                            150761757,
	"TL_CLASS_account_setAccountTTL":                            608323678,
	"TL_CLASS_account_sendChangePhoneCode":                      149257707,
	"TL_CLASS_account_changePhone":                              1891839707,
	"TL_CLASS_account_updateDeviceLocked":                       954152242,
	"TL_CLASS_account_getAuthorizations":                        -484392616,
	"TL_CLASS_account_resetAuthorization":                       -545786948,
	"TL_CLASS_account_getPassword":                              1418342645,
	"TL_CLASS_account_getPasswordSettings":                      -1131605573,
	"TL_CLASS_account_updatePasswordSettings":                   -92517498,
	"TL_CLASS_account_sendConfirmPhoneCode":                     353818557,
	"TL_CLASS_account_confirmPhone":                             1596029123,
	"TL_CLASS_account_getTmpPassword":                           1250046590,
	"TL_CLASS_users_getUsers":                                   227648840,
	"TL_CLASS_users_getFullUser":                                -902781519,
	"TL_CLASS_contacts_getStatuses":                             -995929106,
	"TL_CLASS_contacts_getContacts":                             -1071414113,
	"TL_CLASS_contacts_importContacts":                          746589157,
	"TL_CLASS_contacts_deleteContact":                           -1902823612,
	"TL_CLASS_contacts_deleteContacts":                          1504393374,
	"TL_CLASS_contacts_block":                                   858475004,
	"TL_CLASS_contacts_unblock":                                 -448724803,
	"TL_CLASS_contacts_getBlocked":                              -176409329,
	"TL_CLASS_contacts_exportCard":                              -2065352905,
	"TL_CLASS_contacts_importCard":                              1340184318,
	"TL_CLASS_contacts_search":                                  301470424,
	"TL_CLASS_contacts_resolveUsername":                         -113456221,
	"TL_CLASS_contacts_getTopPeers":                             -728224331,
	"TL_CLASS_contacts_resetTopPeerRating":                      451113900,
	"TL_CLASS_contacts_resetSaved":                              -2020263951,
	"TL_CLASS_messages_getMessages":                             1109588596,
	"TL_CLASS_messages_getDialogs":                              421243333,
	"TL_CLASS_messages_getHistory":                              -591691168,
	"TL_CLASS_messages_search":                                  60726944,
	"TL_CLASS_messages_readHistory":                             238054714,
	"TL_CLASS_messages_deleteHistory":                           469850889,
	"TL_CLASS_messages_deleteMessages":                          -443640366,
	"TL_CLASS_messages_receivedMessages":                        94983360,
	"TL_CLASS_messages_setTyping":                               -1551737264,
	"TL_CLASS_messages_sendMessage":                             -91733382,
	"TL_CLASS_messages_sendMedia":                               -923703407,
	"TL_CLASS_messages_forwardMessages":                         1888354709,
	"TL_CLASS_messages_reportSpam":                              -820669733,
	"TL_CLASS_messages_hideReportSpam":                          -1460572005,
	"TL_CLASS_messages_getPeerSettings":                         913498268,
	"TL_CLASS_messages_getChats":                                1013621127,
	"TL_CLASS_messages_getFullChat":                             998448230,
	"TL_CLASS_messages_editChatTitle":                           -599447467,
	"TL_CLASS_messages_editChatPhoto":                           -900957736,
	"TL_CLASS_messages_addChatUser":                             -106911223,
	"TL_CLASS_messages_deleteChatUser":                          -530505962,
	"TL_CLASS_messages_createChat":                              164303470,
	"TL_CLASS_messages_forwardMessage":                          865483769,
	"TL_CLASS_messages_getDhConfig":                             651135312,
	"TL_CLASS_messages_requestEncryption":                       -162681021,
	"TL_CLASS_messages_acceptEncryption":                        1035731989,
	"TL_CLASS_messages_discardEncryption":                       -304536635,
	"TL_CLASS_messages_setEncryptedTyping":                      2031374829,
	"TL_CLASS_messages_readEncryptedHistory":                    2135648522,
	"TL_CLASS_messages_sendEncrypted":                           -1451792525,
	"TL_CLASS_messages_sendEncryptedFile":                       -1701831834,
	"TL_CLASS_messages_sendEncryptedService":                    852769188,
	"TL_CLASS_messages_receivedQueue":                           1436924774,
	"TL_CLASS_messages_reportEncryptedSpam":                     1259113487,
	"TL_CLASS_messages_readMessageContents":                     916930423,
	"TL_CLASS_messages_getAllStickers":                          479598769,
	"TL_CLASS_messages_getWebPagePreview":                       623001124,
	"TL_CLASS_messages_exportChatInvite":                        2106086025,
	"TL_CLASS_messages_checkChatInvite":                         1051570619,
	"TL_CLASS_messages_importChatInvite":                        1817183516,
	"TL_CLASS_messages_getStickerSet":                           639215886,
	"TL_CLASS_messages_installStickerSet":                       -946871200,
	"TL_CLASS_messages_uninstallStickerSet":                     -110209570,
	"TL_CLASS_messages_startBot":                                -421563528,
	"TL_CLASS_messages_getMessagesViews":                        -993483427,
	"TL_CLASS_messages_toggleChatAdmins":                        -326379039,
	"TL_CLASS_messages_editChatAdmin":                           -1444503762,
	"TL_CLASS_messages_migrateChat":                             363051235,
	"TL_CLASS_messages_searchGlobal":                            -1640190800,
	"TL_CLASS_messages_reorderStickerSets":                      2016638777,
	"TL_CLASS_messages_getDocumentByHash":                       864953444,
	"TL_CLASS_messages_searchGifs":                              -1080395925,
	"TL_CLASS_messages_getSavedGifs":                            -2084618926,
	"TL_CLASS_messages_saveGif":                                 846868683,
	"TL_CLASS_messages_getInlineBotResults":                     1364105629,
	"TL_CLASS_messages_setInlineBotResults":                     -346119674,
	"TL_CLASS_messages_sendInlineBotResult":                     -1318189314,
	"TL_CLASS_messages_getMessageEditData":                      -39416522,
	"TL_CLASS_messages_editMessage":                             97630429,
	"TL_CLASS_messages_editInlineBotMessage":                    -1327463869,
	"TL_CLASS_messages_getBotCallbackAnswer":                    -2130010132,
	"TL_CLASS_messages_setBotCallbackAnswer":                    -712043766,
	"TL_CLASS_messages_getPeerDialogs":                          764901049,
	"TL_CLASS_messages_saveDraft":                               -1137057461,
	"TL_CLASS_messages_getAllDrafts":                            1782549861,
	"TL_CLASS_messages_getFeaturedStickers":                     766298703,
	"TL_CLASS_messages_readFeaturedStickers":                    1527873830,
	"TL_CLASS_messages_getRecentStickers":                       1587647177,
	"TL_CLASS_messages_saveRecentSticker":                       958863608,
	"TL_CLASS_messages_clearRecentStickers":                     -1986437075,
	"TL_CLASS_messages_getArchivedStickers":                     1475442322,
	"TL_CLASS_messages_getMaskStickers":                         1706608543,
	"TL_CLASS_messages_getAttachedStickers":                     -866424884,
	"TL_CLASS_messages_setGameScore":                            -1896289088,
	"TL_CLASS_messages_setInlineGameScore":                      363700068,
	"TL_CLASS_messages_getGameHighScores":                       -400399203,
	"TL_CLASS_messages_getInlineGameHighScores":                 258170395,
	"TL_CLASS_messages_getCommonChats":                          218777796,
	"TL_CLASS_messages_getAllChats":                             -341307408,
	"TL_CLASS_messages_getWebPage":                              852135825,
	"TL_CLASS_messages_toggleDialogPin":                         847887978,
	"TL_CLASS_messages_reorderPinnedDialogs":                    -1784678844,
	"TL_CLASS_messages_getPinnedDialogs":                        -497756594,
	"TL_CLASS_messages_setBotShippingResults":                   -436833542,
	"TL_CLASS_messages_setBotPrecheckoutResults":                163765653,
	"TL_CLASS_messages_uploadMedia":                             1369162417,
	"TL_CLASS_messages_sendScreenshotNotification":              -914493408,
	"TL_CLASS_messages_getFavedStickers":                        567151374,
	"TL_CLASS_messages_faveSticker":                             -1174420133,
	"TL_CLASS_messages_getUnreadMentions":                       1180140658,
	"TL_CLASS_messages_readMentions":                            251759059,
	"TL_CLASS_messages_getRecentLocations":                      613691874,
	"TL_CLASS_messages_sendMultiMedia":                          546656559,
	"TL_CLASS_messages_uploadEncryptedFile":                     1347929239,
	"TL_CLASS_updates_getState":                                 -304838614,
	"TL_CLASS_updates_getDifference":                            630429265,
	"TL_CLASS_updates_getChannelDifference":                     51854712,
	"TL_CLASS_photos_updateProfilePhoto":                        -256159406,
	"TL_CLASS_photos_uploadProfilePhoto":                        1328726168,
	"TL_CLASS_photos_deletePhotos":                              -2016444625,
	"TL_CLASS_photos_getUserPhotos":                             -1848823128,
	"TL_CLASS_upload_saveFilePart":                              -1291540959,
	"TL_CLASS_upload_getFile":                                   -475607115,
	"TL_CLASS_upload_saveBigFilePart":                           -562337987,
	"TL_CLASS_upload_getWebFile":                                619086221,
	"TL_CLASS_upload_getCdnFile":                                536919235,
	"TL_CLASS_upload_reuploadCdnFile":                           452533257,
	"TL_CLASS_upload_getCdnFileHashes":                          -149567365,
	"TL_CLASS_help_getConfig":                                   -990308245,
	"TL_CLASS_help_getNearestDc":                                531836966,
	"TL_CLASS_help_getAppUpdate":                                -1372724842,
	"TL_CLASS_help_saveAppLog":                                  1862465352,
	"TL_CLASS_help_getInviteText":                               1295590211,
	"TL_CLASS_help_getSupport":                                  -1663104819,
	"TL_CLASS_help_getAppChangelog":                             -1877938321,
	"TL_CLASS_help_getTermsOfService":                           889286899,
	"TL_CLASS_help_setBotUpdatesStatus":                         -333262899,
	"TL_CLASS_help_getCdnConfig":                                1375900482,
	"TL_CLASS_help_getRecentMeUrls":                             1036054804,
	"TL_CLASS_channels_readHistory":                             -871347913,
	"TL_CLASS_channels_deleteMessages":                          -2067661490,
	"TL_CLASS_channels_deleteUserHistory":                       -787622117,
	"TL_CLASS_channels_reportSpam":                              -32999408,
	"TL_CLASS_channels_getMessages":                             -1814580409,
	"TL_CLASS_channels_getParticipants":                         306054633,
	"TL_CLASS_channels_getParticipant":                          1416484774,
	"TL_CLASS_channels_getChannels":                             176122811,
	"TL_CLASS_channels_getFullChannel":                          141781513,
	"TL_CLASS_channels_createChannel":                           -192332417,
	"TL_CLASS_channels_editAbout":                               333610782,
	"TL_CLASS_channels_editAdmin":                               548962836,
	"TL_CLASS_channels_editTitle":                               1450044624,
	"TL_CLASS_channels_editPhoto":                               -248621111,
	"TL_CLASS_channels_checkUsername":                           283557164,
	"TL_CLASS_channels_updateUsername":                          890549214,
	"TL_CLASS_channels_joinChannel":                             615851205,
	"TL_CLASS_channels_leaveChannel":                            -130635115,
	"TL_CLASS_channels_inviteToChannel":                         429865580,
	"TL_CLASS_channels_exportInvite":                            -950663035,
	"TL_CLASS_channels_deleteChannel":                           -1072619549,
	"TL_CLASS_channels_toggleInvites":                           1231065863,
	"TL_CLASS_channels_exportMessageLink":                       -934882771,
	"TL_CLASS_channels_toggleSignatures":                        527021574,
	"TL_CLASS_channels_updatePinnedMessage":                     -1490162350,
	"TL_CLASS_channels_getAdminedPublicChannels":                -1920105769,
	"TL_CLASS_channels_editBanned":                              -1076292147,
	"TL_CLASS_channels_getAdminLog":                             870184064,
	"TL_CLASS_channels_setStickers":                             -359881479,
	"TL_CLASS_channels_readMessageContents":                     -357180360,
	"TL_CLASS_channels_deleteHistory":                           -1355375294,
	"TL_CLASS_channels_togglePreHistoryHidden":                  -356796084,
	"TL_CLASS_bots_sendCustomRequest":                           -1440257555,
	"TL_CLASS_bots_answerWebhookJSONQuery":                      -434028723,
	"TL_CLASS_payments_getPaymentForm":                          -1712285883,
	"TL_CLASS_payments_getPaymentReceipt":                       -1601001088,
	"TL_CLASS_payments_validateRequestedInfo":                   1997180532,
	"TL_CLASS_payments_sendPaymentForm":                         730364339,
	"TL_CLASS_payments_getSavedInfo":                            578650699,
	"TL_CLASS_payments_clearSavedInfo":                          -667062079,
	"TL_CLASS_stickers_createStickerSet":                        -1680314774,
	"TL_CLASS_stickers_removeStickerFromSet":                    -143257775,
	"TL_CLASS_stickers_changeStickerPosition":                   -4795190,
	"TL_CLASS_stickers_addStickerToSet":                         -2041315650,
	"TL_CLASS_phone_getCallConfig":                              1430593449,
	"TL_CLASS_phone_requestCall":                                1536537556,
	"TL_CLASS_phone_acceptCall":                                 1003664544,
	"TL_CLASS_phone_confirmCall":                                788404002,
	"TL_CLASS_phone_receivedCall":                               399855457,
	"TL_CLASS_phone_discardCall":                                2027164582,
	"TL_CLASS_phone_setCallRating":                              475228724,
	"TL_CLASS_phone_saveCallDebug":                              662363518,
	"TL_CLASS_langpack_getLangPack":                             -1699363442,
	"TL_CLASS_langpack_getStrings":                              773776152,
	"TL_CLASS_langpack_getDifference":                           187583869,
	"TL_CLASS_langpack_getLanguages":                            -2146445955,
}
