package mtproto

type TL_CLASS_ID int32

const (
	TL_CLASS_UNKNOWN                                          TL_CLASS_ID = 0
	TL_CLASS_resPQ                                            TL_CLASS_ID = 85337187
	TL_CLASS_p_q_inner_data                                   TL_CLASS_ID = -2083955988
	TL_CLASS_server_DH_params_fail                            TL_CLASS_ID = 2043348061
	TL_CLASS_server_DH_params_ok                              TL_CLASS_ID = -790100132
	TL_CLASS_server_DH_inner_data                             TL_CLASS_ID = -1249309254
	TL_CLASS_client_DH_inner_data                             TL_CLASS_ID = 1715713620
	TL_CLASS_dh_gen_ok                                        TL_CLASS_ID = 1003222836
	TL_CLASS_dh_gen_retry                                     TL_CLASS_ID = 1188831161
	TL_CLASS_dh_gen_fail                                      TL_CLASS_ID = -1499615742
	TL_CLASS_destroy_auth_key_ok                              TL_CLASS_ID = -161422892
	TL_CLASS_destroy_auth_key_none                            TL_CLASS_ID = 178201177
	TL_CLASS_destroy_auth_key_fail                            TL_CLASS_ID = -368010477
	TL_CLASS_req_pq                                           TL_CLASS_ID = 1615239032
	TL_CLASS_req_DH_params                                    TL_CLASS_ID = -686627650
	TL_CLASS_set_client_DH_params                             TL_CLASS_ID = -184262881
	TL_CLASS_destroy_auth_key                                 TL_CLASS_ID = -784117408
	TL_CLASS_msgs_ack                                         TL_CLASS_ID = 1658238041
	TL_CLASS_bad_msg_notification                             TL_CLASS_ID = -1477445615
	TL_CLASS_bad_server_salt                                  TL_CLASS_ID = -307542917
	TL_CLASS_msgs_state_req                                   TL_CLASS_ID = -630588590
	TL_CLASS_msgs_state_info                                  TL_CLASS_ID = 81704317
	TL_CLASS_msgs_all_info                                    TL_CLASS_ID = -1933520591
	TL_CLASS_msg_detailed_info                                TL_CLASS_ID = 661470918
	TL_CLASS_msg_new_detailed_info                            TL_CLASS_ID = -2137147681
	TL_CLASS_msg_resend_req                                   TL_CLASS_ID = 2105940488
	TL_CLASS_rpc_error                                        TL_CLASS_ID = 558156313
	TL_CLASS_rpc_answer_unknown                               TL_CLASS_ID = 1579864942
	TL_CLASS_rpc_answer_dropped_running                       TL_CLASS_ID = -847714938
	TL_CLASS_rpc_answer_dropped                               TL_CLASS_ID = -1539647305
	TL_CLASS_future_salt                                      TL_CLASS_ID = 155834844
	TL_CLASS_future_salts                                     TL_CLASS_ID = -1370486635
	TL_CLASS_pong                                             TL_CLASS_ID = 880243653
	TL_CLASS_destroy_session_ok                               TL_CLASS_ID = -501201412
	TL_CLASS_destroy_session_none                             TL_CLASS_ID = 1658015945
	TL_CLASS_new_session_created                              TL_CLASS_ID = -1631450872
	TL_CLASS_http_wait                                        TL_CLASS_ID = -1835453025
	TL_CLASS_ipPort                                           TL_CLASS_ID = -734810765
	TL_CLASS_help_configSimple                                TL_CLASS_ID = -644365371
	TL_CLASS_rpc_drop_answer                                  TL_CLASS_ID = 1491380032
	TL_CLASS_get_future_salts                                 TL_CLASS_ID = -1188971260
	TL_CLASS_ping                                             TL_CLASS_ID = 2059302892
	TL_CLASS_ping_delay_disconnect                            TL_CLASS_ID = -213746804
	TL_CLASS_destroy_session                                  TL_CLASS_ID = -414113498
	TL_CLASS_contest_saveDeveloperInfo                        TL_CLASS_ID = -1705021803
	TL_CLASS_boolFalse                                        TL_CLASS_ID = -1132882121
	TL_CLASS_boolTrue                                         TL_CLASS_ID = -1720552011
	TL_CLASS_true                                             TL_CLASS_ID = 1072550713
	TL_CLASS_vector                                           TL_CLASS_ID = 481674261
	TL_CLASS_error                                            TL_CLASS_ID = -994444869
	TL_CLASS_null                                             TL_CLASS_ID = 1450380236
	TL_CLASS_inputPeerEmpty                                   TL_CLASS_ID = 2134579434
	TL_CLASS_inputPeerSelf                                    TL_CLASS_ID = 2107670217
	TL_CLASS_inputPeerChat                                    TL_CLASS_ID = 396093539
	TL_CLASS_inputPeerUser                                    TL_CLASS_ID = 2072935910
	TL_CLASS_inputPeerChannel                                 TL_CLASS_ID = 548253432
	TL_CLASS_inputUserEmpty                                   TL_CLASS_ID = -1182234929
	TL_CLASS_inputUserSelf                                    TL_CLASS_ID = -138301121
	TL_CLASS_inputUser                                        TL_CLASS_ID = -668391402
	TL_CLASS_inputPhoneContact                                TL_CLASS_ID = -208488460
	TL_CLASS_inputFile                                        TL_CLASS_ID = -181407105
	TL_CLASS_inputFileBig                                     TL_CLASS_ID = -95482955
	TL_CLASS_inputMediaEmpty                                  TL_CLASS_ID = -1771768449
	TL_CLASS_inputMediaUploadedPhoto                          TL_CLASS_ID = 792191537
	TL_CLASS_inputMediaPhoto                                  TL_CLASS_ID = -2114308294
	TL_CLASS_inputMediaGeoPoint                               TL_CLASS_ID = -104578748
	TL_CLASS_inputMediaContact                                TL_CLASS_ID = -1494984313
	TL_CLASS_inputMediaUploadedDocument                       TL_CLASS_ID = -476700163
	TL_CLASS_inputMediaDocument                               TL_CLASS_ID = 1523279502
	TL_CLASS_inputMediaVenue                                  TL_CLASS_ID = -1052959727
	TL_CLASS_inputMediaGifExternal                            TL_CLASS_ID = 1212395773
	TL_CLASS_inputMediaPhotoExternal                          TL_CLASS_ID = 153267905
	TL_CLASS_inputMediaDocumentExternal                       TL_CLASS_ID = -1225309387
	TL_CLASS_inputMediaGame                                   TL_CLASS_ID = -750828557
	TL_CLASS_inputMediaInvoice                                TL_CLASS_ID = -186607933
	TL_CLASS_inputMediaGeoLive                                TL_CLASS_ID = 2065305999
	TL_CLASS_inputChatPhotoEmpty                              TL_CLASS_ID = 480546647
	TL_CLASS_inputChatUploadedPhoto                           TL_CLASS_ID = -1837345356
	TL_CLASS_inputChatPhoto                                   TL_CLASS_ID = -1991004873
	TL_CLASS_inputGeoPointEmpty                               TL_CLASS_ID = -457104426
	TL_CLASS_inputGeoPoint                                    TL_CLASS_ID = -206066487
	TL_CLASS_inputPhotoEmpty                                  TL_CLASS_ID = 483901197
	TL_CLASS_inputPhoto                                       TL_CLASS_ID = -74070332
	TL_CLASS_inputFileLocation                                TL_CLASS_ID = 342061462
	TL_CLASS_inputEncryptedFileLocation                       TL_CLASS_ID = -182231723
	TL_CLASS_inputDocumentFileLocation                        TL_CLASS_ID = 1125058340
	TL_CLASS_inputAppEvent                                    TL_CLASS_ID = 1996904104
	TL_CLASS_peerUser                                         TL_CLASS_ID = -1649296275
	TL_CLASS_peerChat                                         TL_CLASS_ID = -1160714821
	TL_CLASS_peerChannel                                      TL_CLASS_ID = -1109531342
	TL_CLASS_storage_fileUnknown                              TL_CLASS_ID = -1432995067
	TL_CLASS_storage_filePartial                              TL_CLASS_ID = 1086091090
	TL_CLASS_storage_fileJpeg                                 TL_CLASS_ID = 8322574
	TL_CLASS_storage_fileGif                                  TL_CLASS_ID = -891180321
	TL_CLASS_storage_filePng                                  TL_CLASS_ID = 172975040
	TL_CLASS_storage_filePdf                                  TL_CLASS_ID = -1373745011
	TL_CLASS_storage_fileMp3                                  TL_CLASS_ID = 1384777335
	TL_CLASS_storage_fileMov                                  TL_CLASS_ID = 1258941372
	TL_CLASS_storage_fileMp4                                  TL_CLASS_ID = -1278304028
	TL_CLASS_storage_fileWebp                                 TL_CLASS_ID = 276907596
	TL_CLASS_fileLocationUnavailable                          TL_CLASS_ID = 2086234950
	TL_CLASS_fileLocation                                     TL_CLASS_ID = 1406570614
	TL_CLASS_userEmpty                                        TL_CLASS_ID = 537022650
	TL_CLASS_user                                             TL_CLASS_ID = 773059779
	TL_CLASS_userProfilePhotoEmpty                            TL_CLASS_ID = 1326562017
	TL_CLASS_userProfilePhoto                                 TL_CLASS_ID = -715532088
	TL_CLASS_userStatusEmpty                                  TL_CLASS_ID = 164646985
	TL_CLASS_userStatusOnline                                 TL_CLASS_ID = -306628279
	TL_CLASS_userStatusOffline                                TL_CLASS_ID = 9203775
	TL_CLASS_userStatusRecently                               TL_CLASS_ID = -496024847
	TL_CLASS_userStatusLastWeek                               TL_CLASS_ID = 129960444
	TL_CLASS_userStatusLastMonth                              TL_CLASS_ID = 2011940674
	TL_CLASS_chatEmpty                                        TL_CLASS_ID = -1683826688
	TL_CLASS_chat                                             TL_CLASS_ID = -652419756
	TL_CLASS_chatForbidden                                    TL_CLASS_ID = 120753115
	TL_CLASS_channel                                          TL_CLASS_ID = 1158377749
	TL_CLASS_channelForbidden                                 TL_CLASS_ID = 681420594
	TL_CLASS_chatFull                                         TL_CLASS_ID = 771925524
	TL_CLASS_channelFull                                      TL_CLASS_ID = 1991201921
	TL_CLASS_chatParticipant                                  TL_CLASS_ID = -925415106
	TL_CLASS_chatParticipantCreator                           TL_CLASS_ID = -636267638
	TL_CLASS_chatParticipantAdmin                             TL_CLASS_ID = -489233354
	TL_CLASS_chatParticipantsForbidden                        TL_CLASS_ID = -57668565
	TL_CLASS_chatParticipants                                 TL_CLASS_ID = 1061556205
	TL_CLASS_chatPhotoEmpty                                   TL_CLASS_ID = 935395612
	TL_CLASS_chatPhoto                                        TL_CLASS_ID = 1632839530
	TL_CLASS_messageEmpty                                     TL_CLASS_ID = -2082087340
	TL_CLASS_message                                          TL_CLASS_ID = 1157215293
	TL_CLASS_messageService                                   TL_CLASS_ID = -1642487306
	TL_CLASS_messageMediaEmpty                                TL_CLASS_ID = 1038967584
	TL_CLASS_messageMediaPhoto                                TL_CLASS_ID = -1256047857
	TL_CLASS_messageMediaGeo                                  TL_CLASS_ID = 1457575028
	TL_CLASS_messageMediaContact                              TL_CLASS_ID = 1585262393
	TL_CLASS_messageMediaUnsupported                          TL_CLASS_ID = -1618676578
	TL_CLASS_messageMediaDocument                             TL_CLASS_ID = 2084836563
	TL_CLASS_messageMediaWebPage                              TL_CLASS_ID = -1557277184
	TL_CLASS_messageMediaVenue                                TL_CLASS_ID = 784356159
	TL_CLASS_messageMediaGame                                 TL_CLASS_ID = -38694904
	TL_CLASS_messageMediaInvoice                              TL_CLASS_ID = -2074799289
	TL_CLASS_messageMediaGeoLive                              TL_CLASS_ID = 2084316681
	TL_CLASS_messageActionEmpty                               TL_CLASS_ID = -1230047312
	TL_CLASS_messageActionChatCreate                          TL_CLASS_ID = -1503425638
	TL_CLASS_messageActionChatEditTitle                       TL_CLASS_ID = -1247687078
	TL_CLASS_messageActionChatEditPhoto                       TL_CLASS_ID = 2144015272
	TL_CLASS_messageActionChatDeletePhoto                     TL_CLASS_ID = -1780220945
	TL_CLASS_messageActionChatAddUser                         TL_CLASS_ID = 1217033015
	TL_CLASS_messageActionChatDeleteUser                      TL_CLASS_ID = -1297179892
	TL_CLASS_messageActionChatJoinedByLink                    TL_CLASS_ID = -123931160
	TL_CLASS_messageActionChannelCreate                       TL_CLASS_ID = -1781355374
	TL_CLASS_messageActionChatMigrateTo                       TL_CLASS_ID = 1371385889
	TL_CLASS_messageActionChannelMigrateFrom                  TL_CLASS_ID = -1336546578
	TL_CLASS_messageActionPinMessage                          TL_CLASS_ID = -1799538451
	TL_CLASS_messageActionHistoryClear                        TL_CLASS_ID = -1615153660
	TL_CLASS_messageActionGameScore                           TL_CLASS_ID = -1834538890
	TL_CLASS_messageActionPaymentSentMe                       TL_CLASS_ID = -1892568281
	TL_CLASS_messageActionPaymentSent                         TL_CLASS_ID = 1080663248
	TL_CLASS_messageActionPhoneCall                           TL_CLASS_ID = -2132731265
	TL_CLASS_messageActionScreenshotTaken                     TL_CLASS_ID = 1200788123
	TL_CLASS_messageActionCustomAction                        TL_CLASS_ID = -85549226
	TL_CLASS_dialog                                           TL_CLASS_ID = -455150117
	TL_CLASS_photoEmpty                                       TL_CLASS_ID = 590459437
	TL_CLASS_photo                                            TL_CLASS_ID = -1836524247
	TL_CLASS_photoSizeEmpty                                   TL_CLASS_ID = 236446268
	TL_CLASS_photoSize                                        TL_CLASS_ID = 2009052699
	TL_CLASS_photoCachedSize                                  TL_CLASS_ID = -374917894
	TL_CLASS_geoPointEmpty                                    TL_CLASS_ID = 286776671
	TL_CLASS_geoPoint                                         TL_CLASS_ID = 541710092
	TL_CLASS_auth_checkedPhone                                TL_CLASS_ID = -2128698738
	TL_CLASS_auth_sentCode                                    TL_CLASS_ID = 1577067778
	TL_CLASS_auth_authorization                               TL_CLASS_ID = -855308010
	TL_CLASS_auth_exportedAuthorization                       TL_CLASS_ID = -543777747
	TL_CLASS_inputNotifyPeer                                  TL_CLASS_ID = -1195615476
	TL_CLASS_inputNotifyUsers                                 TL_CLASS_ID = 423314455
	TL_CLASS_inputNotifyChats                                 TL_CLASS_ID = 1251338318
	TL_CLASS_inputNotifyAll                                   TL_CLASS_ID = -1540769658
	TL_CLASS_inputPeerNotifyEventsEmpty                       TL_CLASS_ID = -265263912
	TL_CLASS_inputPeerNotifyEventsAll                         TL_CLASS_ID = -395694988
	TL_CLASS_inputPeerNotifySettings                          TL_CLASS_ID = 949182130
	TL_CLASS_peerNotifyEventsEmpty                            TL_CLASS_ID = -1378534221
	TL_CLASS_peerNotifyEventsAll                              TL_CLASS_ID = 1830677896
	TL_CLASS_peerNotifySettingsEmpty                          TL_CLASS_ID = 1889961234
	TL_CLASS_peerNotifySettings                               TL_CLASS_ID = -1697798976
	TL_CLASS_peerSettings                                     TL_CLASS_ID = -2122045747
	TL_CLASS_wallPaper                                        TL_CLASS_ID = -860866985
	TL_CLASS_wallPaperSolid                                   TL_CLASS_ID = 1662091044
	TL_CLASS_inputReportReasonSpam                            TL_CLASS_ID = 1490799288
	TL_CLASS_inputReportReasonViolence                        TL_CLASS_ID = 505595789
	TL_CLASS_inputReportReasonPornography                     TL_CLASS_ID = 777640226
	TL_CLASS_inputReportReasonOther                           TL_CLASS_ID = -512463606
	TL_CLASS_userFull                                         TL_CLASS_ID = 253890367
	TL_CLASS_contact                                          TL_CLASS_ID = -116274796
	TL_CLASS_importedContact                                  TL_CLASS_ID = -805141448
	TL_CLASS_contactBlocked                                   TL_CLASS_ID = 1444661369
	TL_CLASS_contactStatus                                    TL_CLASS_ID = -748155807
	TL_CLASS_contacts_link                                    TL_CLASS_ID = 986597452
	TL_CLASS_contacts_contactsNotModified                     TL_CLASS_ID = -1219778094
	TL_CLASS_contacts_contacts                                TL_CLASS_ID = -353862078
	TL_CLASS_contacts_importedContacts                        TL_CLASS_ID = 2010127419
	TL_CLASS_contacts_blocked                                 TL_CLASS_ID = 471043349
	TL_CLASS_contacts_blockedSlice                            TL_CLASS_ID = -1878523231
	TL_CLASS_messages_dialogs                                 TL_CLASS_ID = 364538944
	TL_CLASS_messages_dialogsSlice                            TL_CLASS_ID = 1910543603
	TL_CLASS_messages_messages                                TL_CLASS_ID = -1938715001
	TL_CLASS_messages_messagesSlice                           TL_CLASS_ID = 189033187
	TL_CLASS_messages_channelMessages                         TL_CLASS_ID = -1725551049
	TL_CLASS_messages_messagesNotModified                     TL_CLASS_ID = 1951620897
	TL_CLASS_messages_chats                                   TL_CLASS_ID = 1694474197
	TL_CLASS_messages_chatsSlice                              TL_CLASS_ID = -1663561404
	TL_CLASS_messages_chatFull                                TL_CLASS_ID = -438840932
	TL_CLASS_messages_affectedHistory                         TL_CLASS_ID = -1269012015
	TL_CLASS_inputMessagesFilterEmpty                         TL_CLASS_ID = 1474492012
	TL_CLASS_inputMessagesFilterPhotos                        TL_CLASS_ID = -1777752804
	TL_CLASS_inputMessagesFilterVideo                         TL_CLASS_ID = -1614803355
	TL_CLASS_inputMessagesFilterPhotoVideo                    TL_CLASS_ID = 1458172132
	TL_CLASS_inputMessagesFilterDocument                      TL_CLASS_ID = -1629621880
	TL_CLASS_inputMessagesFilterUrl                           TL_CLASS_ID = 2129714567
	TL_CLASS_inputMessagesFilterGif                           TL_CLASS_ID = -3644025
	TL_CLASS_inputMessagesFilterVoice                         TL_CLASS_ID = 1358283666
	TL_CLASS_inputMessagesFilterMusic                         TL_CLASS_ID = 928101534
	TL_CLASS_inputMessagesFilterChatPhotos                    TL_CLASS_ID = 975236280
	TL_CLASS_inputMessagesFilterPhoneCalls                    TL_CLASS_ID = -2134272152
	TL_CLASS_inputMessagesFilterRoundVoice                    TL_CLASS_ID = 2054952868
	TL_CLASS_inputMessagesFilterRoundVideo                    TL_CLASS_ID = -1253451181
	TL_CLASS_inputMessagesFilterMyMentions                    TL_CLASS_ID = -1040652646
	TL_CLASS_inputMessagesFilterGeo                           TL_CLASS_ID = -419271411
	TL_CLASS_inputMessagesFilterContacts                      TL_CLASS_ID = -530392189
	TL_CLASS_updateNewMessage                                 TL_CLASS_ID = 522914557
	TL_CLASS_updateMessageID                                  TL_CLASS_ID = 1318109142
	TL_CLASS_updateDeleteMessages                             TL_CLASS_ID = -1576161051
	TL_CLASS_updateUserTyping                                 TL_CLASS_ID = 1548249383
	TL_CLASS_updateChatUserTyping                             TL_CLASS_ID = -1704596961
	TL_CLASS_updateChatParticipants                           TL_CLASS_ID = 125178264
	TL_CLASS_updateUserStatus                                 TL_CLASS_ID = 469489699
	TL_CLASS_updateUserName                                   TL_CLASS_ID = -1489818765
	TL_CLASS_updateUserPhoto                                  TL_CLASS_ID = -1791935732
	TL_CLASS_updateContactRegistered                          TL_CLASS_ID = 628472761
	TL_CLASS_updateContactLink                                TL_CLASS_ID = -1657903163
	TL_CLASS_updateNewEncryptedMessage                        TL_CLASS_ID = 314359194
	TL_CLASS_updateEncryptedChatTyping                        TL_CLASS_ID = 386986326
	TL_CLASS_updateEncryption                                 TL_CLASS_ID = -1264392051
	TL_CLASS_updateEncryptedMessagesRead                      TL_CLASS_ID = 956179895
	TL_CLASS_updateChatParticipantAdd                         TL_CLASS_ID = -364179876
	TL_CLASS_updateChatParticipantDelete                      TL_CLASS_ID = 1851755554
	TL_CLASS_updateDcOptions                                  TL_CLASS_ID = -1906403213
	TL_CLASS_updateUserBlocked                                TL_CLASS_ID = -2131957734
	TL_CLASS_updateNotifySettings                             TL_CLASS_ID = -1094555409
	TL_CLASS_updateServiceNotification                        TL_CLASS_ID = -337352679
	TL_CLASS_updatePrivacy                                    TL_CLASS_ID = -298113238
	TL_CLASS_updateUserPhone                                  TL_CLASS_ID = 314130811
	TL_CLASS_updateReadHistoryInbox                           TL_CLASS_ID = -1721631396
	TL_CLASS_updateReadHistoryOutbox                          TL_CLASS_ID = 791617983
	TL_CLASS_updateWebPage                                    TL_CLASS_ID = 2139689491
	TL_CLASS_updateReadMessagesContents                       TL_CLASS_ID = 1757493555
	TL_CLASS_updateChannelTooLong                             TL_CLASS_ID = -352032773
	TL_CLASS_updateChannel                                    TL_CLASS_ID = -1227598250
	TL_CLASS_updateNewChannelMessage                          TL_CLASS_ID = 1656358105
	TL_CLASS_updateReadChannelInbox                           TL_CLASS_ID = 1108669311
	TL_CLASS_updateDeleteChannelMessages                      TL_CLASS_ID = -1015733815
	TL_CLASS_updateChannelMessageViews                        TL_CLASS_ID = -1734268085
	TL_CLASS_updateChatAdmins                                 TL_CLASS_ID = 1855224129
	TL_CLASS_updateChatParticipantAdmin                       TL_CLASS_ID = -1232070311
	TL_CLASS_updateNewStickerSet                              TL_CLASS_ID = 1753886890
	TL_CLASS_updateStickerSetsOrder                           TL_CLASS_ID = 196268545
	TL_CLASS_updateStickerSets                                TL_CLASS_ID = 1135492588
	TL_CLASS_updateSavedGifs                                  TL_CLASS_ID = -1821035490
	TL_CLASS_updateBotInlineQuery                             TL_CLASS_ID = 1417832080
	TL_CLASS_updateBotInlineSend                              TL_CLASS_ID = 239663460
	TL_CLASS_updateEditChannelMessage                         TL_CLASS_ID = 457133559
	TL_CLASS_updateChannelPinnedMessage                       TL_CLASS_ID = -1738988427
	TL_CLASS_updateBotCallbackQuery                           TL_CLASS_ID = -415938591
	TL_CLASS_updateEditMessage                                TL_CLASS_ID = -469536605
	TL_CLASS_updateInlineBotCallbackQuery                     TL_CLASS_ID = -103646630
	TL_CLASS_updateReadChannelOutbox                          TL_CLASS_ID = 634833351
	TL_CLASS_updateDraftMessage                               TL_CLASS_ID = -299124375
	TL_CLASS_updateReadFeaturedStickers                       TL_CLASS_ID = 1461528386
	TL_CLASS_updateRecentStickers                             TL_CLASS_ID = -1706939360
	TL_CLASS_updateConfig                                     TL_CLASS_ID = -1574314746
	TL_CLASS_updatePtsChanged                                 TL_CLASS_ID = 861169551
	TL_CLASS_updateChannelWebPage                             TL_CLASS_ID = 1081547008
	TL_CLASS_updateDialogPinned                               TL_CLASS_ID = -686710068
	TL_CLASS_updatePinnedDialogs                              TL_CLASS_ID = -657787251
	TL_CLASS_updateBotWebhookJSON                             TL_CLASS_ID = -2095595325
	TL_CLASS_updateBotWebhookJSONQuery                        TL_CLASS_ID = -1684914010
	TL_CLASS_updateBotShippingQuery                           TL_CLASS_ID = -523384512
	TL_CLASS_updateBotPrecheckoutQuery                        TL_CLASS_ID = 1563376297
	TL_CLASS_updatePhoneCall                                  TL_CLASS_ID = -1425052898
	TL_CLASS_updateLangPackTooLong                            TL_CLASS_ID = 281165899
	TL_CLASS_updateLangPack                                   TL_CLASS_ID = 1442983757
	TL_CLASS_updateFavedStickers                              TL_CLASS_ID = -451831443
	TL_CLASS_updateChannelReadMessagesContents                TL_CLASS_ID = -1987495099
	TL_CLASS_updateContactsReset                              TL_CLASS_ID = 1887741886
	TL_CLASS_updateChannelAvailableMessages                   TL_CLASS_ID = 1893427255
	TL_CLASS_updates_state                                    TL_CLASS_ID = -1519637954
	TL_CLASS_updates_differenceEmpty                          TL_CLASS_ID = 1567990072
	TL_CLASS_updates_difference                               TL_CLASS_ID = 16030880
	TL_CLASS_updates_differenceSlice                          TL_CLASS_ID = -1459938943
	TL_CLASS_updates_differenceTooLong                        TL_CLASS_ID = 1258196845
	TL_CLASS_updatesTooLong                                   TL_CLASS_ID = -484987010
	TL_CLASS_updateShortMessage                               TL_CLASS_ID = -1857044719
	TL_CLASS_updateShortChatMessage                           TL_CLASS_ID = 377562760
	TL_CLASS_updateShort                                      TL_CLASS_ID = 2027216577
	TL_CLASS_updatesCombined                                  TL_CLASS_ID = 1918567619
	TL_CLASS_updates                                          TL_CLASS_ID = 1957577280
	TL_CLASS_updateShortSentMessage                           TL_CLASS_ID = 301019932
	TL_CLASS_photos_photos                                    TL_CLASS_ID = -1916114267
	TL_CLASS_photos_photosSlice                               TL_CLASS_ID = 352657236
	TL_CLASS_photos_photo                                     TL_CLASS_ID = 539045032
	TL_CLASS_upload_file                                      TL_CLASS_ID = 157948117
	TL_CLASS_upload_fileCdnRedirect                           TL_CLASS_ID = -363659686
	TL_CLASS_dcOption                                         TL_CLASS_ID = 98092748
	TL_CLASS_config                                           TL_CLASS_ID = -1669068444
	TL_CLASS_nearestDc                                        TL_CLASS_ID = -1910892683
	TL_CLASS_help_appUpdate                                   TL_CLASS_ID = -1987579119
	TL_CLASS_help_noAppUpdate                                 TL_CLASS_ID = -1000708810
	TL_CLASS_help_inviteText                                  TL_CLASS_ID = 415997816
	TL_CLASS_encryptedChatEmpty                               TL_CLASS_ID = -1417756512
	TL_CLASS_encryptedChatWaiting                             TL_CLASS_ID = 1006044124
	TL_CLASS_encryptedChatRequested                           TL_CLASS_ID = -931638658
	TL_CLASS_encryptedChat                                    TL_CLASS_ID = -94974410
	TL_CLASS_encryptedChatDiscarded                           TL_CLASS_ID = 332848423
	TL_CLASS_inputEncryptedChat                               TL_CLASS_ID = -247351839
	TL_CLASS_encryptedFileEmpty                               TL_CLASS_ID = -1038136962
	TL_CLASS_encryptedFile                                    TL_CLASS_ID = 1248893260
	TL_CLASS_inputEncryptedFileEmpty                          TL_CLASS_ID = 406307684
	TL_CLASS_inputEncryptedFileUploaded                       TL_CLASS_ID = 1690108678
	TL_CLASS_inputEncryptedFile                               TL_CLASS_ID = 1511503333
	TL_CLASS_inputEncryptedFileBigUploaded                    TL_CLASS_ID = 767652808
	TL_CLASS_encryptedMessage                                 TL_CLASS_ID = -317144808
	TL_CLASS_encryptedMessageService                          TL_CLASS_ID = 594758406
	TL_CLASS_messages_dhConfigNotModified                     TL_CLASS_ID = -1058912715
	TL_CLASS_messages_dhConfig                                TL_CLASS_ID = 740433629
	TL_CLASS_messages_sentEncryptedMessage                    TL_CLASS_ID = 1443858741
	TL_CLASS_messages_sentEncryptedFile                       TL_CLASS_ID = -1802240206
	TL_CLASS_inputDocumentEmpty                               TL_CLASS_ID = 1928391342
	TL_CLASS_inputDocument                                    TL_CLASS_ID = 410618194
	TL_CLASS_documentEmpty                                    TL_CLASS_ID = 922273905
	TL_CLASS_document                                         TL_CLASS_ID = -2027738169
	TL_CLASS_help_support                                     TL_CLASS_ID = 398898678
	TL_CLASS_notifyPeer                                       TL_CLASS_ID = -1613493288
	TL_CLASS_notifyUsers                                      TL_CLASS_ID = -1261946036
	TL_CLASS_notifyChats                                      TL_CLASS_ID = -1073230141
	TL_CLASS_notifyAll                                        TL_CLASS_ID = 1959820384
	TL_CLASS_sendMessageTypingAction                          TL_CLASS_ID = 381645902
	TL_CLASS_sendMessageCancelAction                          TL_CLASS_ID = -44119819
	TL_CLASS_sendMessageRecordVideoAction                     TL_CLASS_ID = -1584933265
	TL_CLASS_sendMessageUploadVideoAction                     TL_CLASS_ID = -378127636
	TL_CLASS_sendMessageRecordAudioAction                     TL_CLASS_ID = -718310409
	TL_CLASS_sendMessageUploadAudioAction                     TL_CLASS_ID = -212740181
	TL_CLASS_sendMessageUploadPhotoAction                     TL_CLASS_ID = -774682074
	TL_CLASS_sendMessageUploadDocumentAction                  TL_CLASS_ID = -1441998364
	TL_CLASS_sendMessageGeoLocationAction                     TL_CLASS_ID = 393186209
	TL_CLASS_sendMessageChooseContactAction                   TL_CLASS_ID = 1653390447
	TL_CLASS_sendMessageGamePlayAction                        TL_CLASS_ID = -580219064
	TL_CLASS_sendMessageRecordRoundAction                     TL_CLASS_ID = -1997373508
	TL_CLASS_sendMessageUploadRoundAction                     TL_CLASS_ID = 608050278
	TL_CLASS_contacts_found                                   TL_CLASS_ID = 446822276
	TL_CLASS_inputPrivacyKeyStatusTimestamp                   TL_CLASS_ID = 1335282456
	TL_CLASS_inputPrivacyKeyChatInvite                        TL_CLASS_ID = -1107622874
	TL_CLASS_inputPrivacyKeyPhoneCall                         TL_CLASS_ID = -88417185
	TL_CLASS_privacyKeyStatusTimestamp                        TL_CLASS_ID = -1137792208
	TL_CLASS_privacyKeyChatInvite                             TL_CLASS_ID = 1343122938
	TL_CLASS_privacyKeyPhoneCall                              TL_CLASS_ID = 1030105979
	TL_CLASS_inputPrivacyValueAllowContacts                   TL_CLASS_ID = 218751099
	TL_CLASS_inputPrivacyValueAllowAll                        TL_CLASS_ID = 407582158
	TL_CLASS_inputPrivacyValueAllowUsers                      TL_CLASS_ID = 320652927
	TL_CLASS_inputPrivacyValueDisallowContacts                TL_CLASS_ID = 195371015
	TL_CLASS_inputPrivacyValueDisallowAll                     TL_CLASS_ID = -697604407
	TL_CLASS_inputPrivacyValueDisallowUsers                   TL_CLASS_ID = -1877932953
	TL_CLASS_privacyValueAllowContacts                        TL_CLASS_ID = -123988
	TL_CLASS_privacyValueAllowAll                             TL_CLASS_ID = 1698855810
	TL_CLASS_privacyValueAllowUsers                           TL_CLASS_ID = 1297858060
	TL_CLASS_privacyValueDisallowContacts                     TL_CLASS_ID = -125240806
	TL_CLASS_privacyValueDisallowAll                          TL_CLASS_ID = -1955338397
	TL_CLASS_privacyValueDisallowUsers                        TL_CLASS_ID = 209668535
	TL_CLASS_account_privacyRules                             TL_CLASS_ID = 1430961007
	TL_CLASS_accountDaysTTL                                   TL_CLASS_ID = -1194283041
	TL_CLASS_documentAttributeImageSize                       TL_CLASS_ID = 1815593308
	TL_CLASS_documentAttributeAnimated                        TL_CLASS_ID = 297109817
	TL_CLASS_documentAttributeSticker                         TL_CLASS_ID = 1662637586
	TL_CLASS_documentAttributeVideo                           TL_CLASS_ID = 250621158
	TL_CLASS_documentAttributeAudio                           TL_CLASS_ID = -1739392570
	TL_CLASS_documentAttributeFilename                        TL_CLASS_ID = 358154344
	TL_CLASS_documentAttributeHasStickers                     TL_CLASS_ID = -1744710921
	TL_CLASS_messages_stickersNotModified                     TL_CLASS_ID = -244016606
	TL_CLASS_messages_stickers                                TL_CLASS_ID = -1970352846
	TL_CLASS_stickerPack                                      TL_CLASS_ID = 313694676
	TL_CLASS_messages_allStickersNotModified                  TL_CLASS_ID = -395967805
	TL_CLASS_messages_allStickers                             TL_CLASS_ID = -302170017
	TL_CLASS_disabledFeature                                  TL_CLASS_ID = -1369215196
	TL_CLASS_messages_affectedMessages                        TL_CLASS_ID = -2066640507
	TL_CLASS_contactLinkUnknown                               TL_CLASS_ID = 1599050311
	TL_CLASS_contactLinkNone                                  TL_CLASS_ID = -17968211
	TL_CLASS_contactLinkHasPhone                              TL_CLASS_ID = 646922073
	TL_CLASS_contactLinkContact                               TL_CLASS_ID = -721239344
	TL_CLASS_webPageEmpty                                     TL_CLASS_ID = -350980120
	TL_CLASS_webPagePending                                   TL_CLASS_ID = -981018084
	TL_CLASS_webPage                                          TL_CLASS_ID = 1594340540
	TL_CLASS_webPageNotModified                               TL_CLASS_ID = -2054908813
	TL_CLASS_authorization                                    TL_CLASS_ID = 2079516406
	TL_CLASS_account_authorizations                           TL_CLASS_ID = 307276766
	TL_CLASS_account_noPassword                               TL_CLASS_ID = -1764049896
	TL_CLASS_account_password                                 TL_CLASS_ID = 2081952796
	TL_CLASS_account_passwordSettings                         TL_CLASS_ID = -1212732749
	TL_CLASS_account_passwordInputSettings                    TL_CLASS_ID = -2037289493
	TL_CLASS_auth_passwordRecovery                            TL_CLASS_ID = 326715557
	TL_CLASS_receivedNotifyMessage                            TL_CLASS_ID = -1551583367
	TL_CLASS_chatInviteEmpty                                  TL_CLASS_ID = 1776236393
	TL_CLASS_chatInviteExported                               TL_CLASS_ID = -64092740
	TL_CLASS_chatInviteAlready                                TL_CLASS_ID = 1516793212
	TL_CLASS_chatInvite                                       TL_CLASS_ID = -613092008
	TL_CLASS_inputStickerSetEmpty                             TL_CLASS_ID = -4838507
	TL_CLASS_inputStickerSetID                                TL_CLASS_ID = -1645763991
	TL_CLASS_inputStickerSetShortName                         TL_CLASS_ID = -2044933984
	TL_CLASS_stickerSet                                       TL_CLASS_ID = -852477119
	TL_CLASS_messages_stickerSet                              TL_CLASS_ID = -1240849242
	TL_CLASS_botCommand                                       TL_CLASS_ID = -1032140601
	TL_CLASS_botInfo                                          TL_CLASS_ID = -1729618630
	TL_CLASS_keyboardButton                                   TL_CLASS_ID = -1560655744
	TL_CLASS_keyboardButtonUrl                                TL_CLASS_ID = 629866245
	TL_CLASS_keyboardButtonCallback                           TL_CLASS_ID = 1748655686
	TL_CLASS_keyboardButtonRequestPhone                       TL_CLASS_ID = -1318425559
	TL_CLASS_keyboardButtonRequestGeoLocation                 TL_CLASS_ID = -59151553
	TL_CLASS_keyboardButtonSwitchInline                       TL_CLASS_ID = 90744648
	TL_CLASS_keyboardButtonGame                               TL_CLASS_ID = 1358175439
	TL_CLASS_keyboardButtonBuy                                TL_CLASS_ID = -1344716869
	TL_CLASS_keyboardButtonRow                                TL_CLASS_ID = 2002815875
	TL_CLASS_replyKeyboardHide                                TL_CLASS_ID = -1606526075
	TL_CLASS_replyKeyboardForceReply                          TL_CLASS_ID = -200242528
	TL_CLASS_replyKeyboardMarkup                              TL_CLASS_ID = 889353612
	TL_CLASS_replyInlineMarkup                                TL_CLASS_ID = 1218642516
	TL_CLASS_messageEntityUnknown                             TL_CLASS_ID = -1148011883
	TL_CLASS_messageEntityMention                             TL_CLASS_ID = -100378723
	TL_CLASS_messageEntityHashtag                             TL_CLASS_ID = 1868782349
	TL_CLASS_messageEntityBotCommand                          TL_CLASS_ID = 1827637959
	TL_CLASS_messageEntityUrl                                 TL_CLASS_ID = 1859134776
	TL_CLASS_messageEntityEmail                               TL_CLASS_ID = 1692693954
	TL_CLASS_messageEntityBold                                TL_CLASS_ID = -1117713463
	TL_CLASS_messageEntityItalic                              TL_CLASS_ID = -2106619040
	TL_CLASS_messageEntityCode                                TL_CLASS_ID = 681706865
	TL_CLASS_messageEntityPre                                 TL_CLASS_ID = 1938967520
	TL_CLASS_messageEntityTextUrl                             TL_CLASS_ID = 1990644519
	TL_CLASS_messageEntityMentionName                         TL_CLASS_ID = 892193368
	TL_CLASS_inputMessageEntityMentionName                    TL_CLASS_ID = 546203849
	TL_CLASS_inputChannelEmpty                                TL_CLASS_ID = -292807034
	TL_CLASS_inputChannel                                     TL_CLASS_ID = -1343524562
	TL_CLASS_contacts_resolvedPeer                            TL_CLASS_ID = 2131196633
	TL_CLASS_messageRange                                     TL_CLASS_ID = 182649427
	TL_CLASS_updates_channelDifferenceEmpty                   TL_CLASS_ID = 1041346555
	TL_CLASS_updates_channelDifferenceTooLong                 TL_CLASS_ID = 1788705589
	TL_CLASS_updates_channelDifference                        TL_CLASS_ID = 543450958
	TL_CLASS_channelMessagesFilterEmpty                       TL_CLASS_ID = -1798033689
	TL_CLASS_channelMessagesFilter                            TL_CLASS_ID = -847783593
	TL_CLASS_channelParticipant                               TL_CLASS_ID = 367766557
	TL_CLASS_channelParticipantSelf                           TL_CLASS_ID = -1557620115
	TL_CLASS_channelParticipantCreator                        TL_CLASS_ID = -471670279
	TL_CLASS_channelParticipantAdmin                          TL_CLASS_ID = -1473271656
	TL_CLASS_channelParticipantBanned                         TL_CLASS_ID = 573315206
	TL_CLASS_channelParticipantsRecent                        TL_CLASS_ID = -566281095
	TL_CLASS_channelParticipantsAdmins                        TL_CLASS_ID = -1268741783
	TL_CLASS_channelParticipantsKicked                        TL_CLASS_ID = -1548400251
	TL_CLASS_channelParticipantsBots                          TL_CLASS_ID = -1328445861
	TL_CLASS_channelParticipantsBanned                        TL_CLASS_ID = 338142689
	TL_CLASS_channelParticipantsSearch                        TL_CLASS_ID = 106343499
	TL_CLASS_channels_channelParticipants                     TL_CLASS_ID = -177282392
	TL_CLASS_channels_channelParticipantsNotModified          TL_CLASS_ID = -266911767
	TL_CLASS_channels_channelParticipant                      TL_CLASS_ID = -791039645
	TL_CLASS_help_termsOfService                              TL_CLASS_ID = -236044656
	TL_CLASS_foundGif                                         TL_CLASS_ID = 372165663
	TL_CLASS_foundGifCached                                   TL_CLASS_ID = -1670052855
	TL_CLASS_messages_foundGifs                               TL_CLASS_ID = 1158290442
	TL_CLASS_messages_savedGifsNotModified                    TL_CLASS_ID = -402498398
	TL_CLASS_messages_savedGifs                               TL_CLASS_ID = 772213157
	TL_CLASS_inputBotInlineMessageMediaAuto                   TL_CLASS_ID = 691006739
	TL_CLASS_inputBotInlineMessageText                        TL_CLASS_ID = 1036876423
	TL_CLASS_inputBotInlineMessageMediaGeo                    TL_CLASS_ID = -1045340827
	TL_CLASS_inputBotInlineMessageMediaVenue                  TL_CLASS_ID = -1431327288
	TL_CLASS_inputBotInlineMessageMediaContact                TL_CLASS_ID = 766443943
	TL_CLASS_inputBotInlineMessageGame                        TL_CLASS_ID = 1262639204
	TL_CLASS_inputBotInlineResult                             TL_CLASS_ID = 750510426
	TL_CLASS_inputBotInlineResultPhoto                        TL_CLASS_ID = -1462213465
	TL_CLASS_inputBotInlineResultDocument                     TL_CLASS_ID = -459324
	TL_CLASS_inputBotInlineResultGame                         TL_CLASS_ID = 1336154098
	TL_CLASS_botInlineMessageMediaAuto                        TL_CLASS_ID = 175419739
	TL_CLASS_botInlineMessageText                             TL_CLASS_ID = -1937807902
	TL_CLASS_botInlineMessageMediaGeo                         TL_CLASS_ID = -1222451611
	TL_CLASS_botInlineMessageMediaVenue                       TL_CLASS_ID = 1130767150
	TL_CLASS_botInlineMessageMediaContact                     TL_CLASS_ID = 904770772
	TL_CLASS_botInlineResult                                  TL_CLASS_ID = -1679053127
	TL_CLASS_botInlineMediaResult                             TL_CLASS_ID = 400266251
	TL_CLASS_messages_botResults                              TL_CLASS_ID = -1803769784
	TL_CLASS_exportedMessageLink                              TL_CLASS_ID = 524838915
	TL_CLASS_messageFwdHeader                                 TL_CLASS_ID = 1436466797
	TL_CLASS_auth_codeTypeSms                                 TL_CLASS_ID = 1923290508
	TL_CLASS_auth_codeTypeCall                                TL_CLASS_ID = 1948046307
	TL_CLASS_auth_codeTypeFlashCall                           TL_CLASS_ID = 577556219
	TL_CLASS_auth_sentCodeTypeApp                             TL_CLASS_ID = 1035688326
	TL_CLASS_auth_sentCodeTypeSms                             TL_CLASS_ID = -1073693790
	TL_CLASS_auth_sentCodeTypeCall                            TL_CLASS_ID = 1398007207
	TL_CLASS_auth_sentCodeTypeFlashCall                       TL_CLASS_ID = -1425815847
	TL_CLASS_messages_botCallbackAnswer                       TL_CLASS_ID = 911761060
	TL_CLASS_messages_messageEditData                         TL_CLASS_ID = 649453030
	TL_CLASS_inputBotInlineMessageID                          TL_CLASS_ID = -1995686519
	TL_CLASS_inlineBotSwitchPM                                TL_CLASS_ID = 1008755359
	TL_CLASS_messages_peerDialogs                             TL_CLASS_ID = 863093588
	TL_CLASS_topPeer                                          TL_CLASS_ID = -305282981
	TL_CLASS_topPeerCategoryBotsPM                            TL_CLASS_ID = -1419371685
	TL_CLASS_topPeerCategoryBotsInline                        TL_CLASS_ID = 344356834
	TL_CLASS_topPeerCategoryCorrespondents                    TL_CLASS_ID = 104314861
	TL_CLASS_topPeerCategoryGroups                            TL_CLASS_ID = -1122524854
	TL_CLASS_topPeerCategoryChannels                          TL_CLASS_ID = 371037736
	TL_CLASS_topPeerCategoryPhoneCalls                        TL_CLASS_ID = 511092620
	TL_CLASS_topPeerCategoryPeers                             TL_CLASS_ID = -75283823
	TL_CLASS_contacts_topPeersNotModified                     TL_CLASS_ID = -567906571
	TL_CLASS_contacts_topPeers                                TL_CLASS_ID = 1891070632
	TL_CLASS_draftMessageEmpty                                TL_CLASS_ID = -1169445179
	TL_CLASS_draftMessage                                     TL_CLASS_ID = -40996577
	TL_CLASS_messages_featuredStickersNotModified             TL_CLASS_ID = 82699215
	TL_CLASS_messages_featuredStickers                        TL_CLASS_ID = -123893531
	TL_CLASS_messages_recentStickersNotModified               TL_CLASS_ID = 186120336
	TL_CLASS_messages_recentStickers                          TL_CLASS_ID = 1558317424
	TL_CLASS_messages_archivedStickers                        TL_CLASS_ID = 1338747336
	TL_CLASS_messages_stickerSetInstallResultSuccess          TL_CLASS_ID = 946083368
	TL_CLASS_messages_stickerSetInstallResultArchive          TL_CLASS_ID = 904138920
	TL_CLASS_stickerSetCovered                                TL_CLASS_ID = 1678812626
	TL_CLASS_stickerSetMultiCovered                           TL_CLASS_ID = 872932635
	TL_CLASS_maskCoords                                       TL_CLASS_ID = -1361650766
	TL_CLASS_inputStickeredMediaPhoto                         TL_CLASS_ID = 1251549527
	TL_CLASS_inputStickeredMediaDocument                      TL_CLASS_ID = 70813275
	TL_CLASS_game                                             TL_CLASS_ID = -1107729093
	TL_CLASS_inputGameID                                      TL_CLASS_ID = 53231223
	TL_CLASS_inputGameShortName                               TL_CLASS_ID = -1020139510
	TL_CLASS_highScore                                        TL_CLASS_ID = 1493171408
	TL_CLASS_messages_highScores                              TL_CLASS_ID = -1707344487
	TL_CLASS_textEmpty                                        TL_CLASS_ID = -599948721
	TL_CLASS_textPlain                                        TL_CLASS_ID = 1950782688
	TL_CLASS_textBold                                         TL_CLASS_ID = 1730456516
	TL_CLASS_textItalic                                       TL_CLASS_ID = -653089380
	TL_CLASS_textUnderline                                    TL_CLASS_ID = -1054465340
	TL_CLASS_textStrike                                       TL_CLASS_ID = -1678197867
	TL_CLASS_textFixed                                        TL_CLASS_ID = 1816074681
	TL_CLASS_textUrl                                          TL_CLASS_ID = 1009288385
	TL_CLASS_textEmail                                        TL_CLASS_ID = -564523562
	TL_CLASS_textConcat                                       TL_CLASS_ID = 2120376535
	TL_CLASS_pageBlockUnsupported                             TL_CLASS_ID = 324435594
	TL_CLASS_pageBlockTitle                                   TL_CLASS_ID = 1890305021
	TL_CLASS_pageBlockSubtitle                                TL_CLASS_ID = -1879401953
	TL_CLASS_pageBlockAuthorDate                              TL_CLASS_ID = -1162877472
	TL_CLASS_pageBlockHeader                                  TL_CLASS_ID = -1076861716
	TL_CLASS_pageBlockSubheader                               TL_CLASS_ID = -248793375
	TL_CLASS_pageBlockParagraph                               TL_CLASS_ID = 1182402406
	TL_CLASS_pageBlockPreformatted                            TL_CLASS_ID = -1066346178
	TL_CLASS_pageBlockFooter                                  TL_CLASS_ID = 1216809369
	TL_CLASS_pageBlockDivider                                 TL_CLASS_ID = -618614392
	TL_CLASS_pageBlockAnchor                                  TL_CLASS_ID = -837994576
	TL_CLASS_pageBlockList                                    TL_CLASS_ID = 978896884
	TL_CLASS_pageBlockBlockquote                              TL_CLASS_ID = 641563686
	TL_CLASS_pageBlockPullquote                               TL_CLASS_ID = 1329878739
	TL_CLASS_pageBlockPhoto                                   TL_CLASS_ID = -372860542
	TL_CLASS_pageBlockVideo                                   TL_CLASS_ID = -640214938
	TL_CLASS_pageBlockCover                                   TL_CLASS_ID = 972174080
	TL_CLASS_pageBlockEmbed                                   TL_CLASS_ID = -840826671
	TL_CLASS_pageBlockEmbedPost                               TL_CLASS_ID = 690781161
	TL_CLASS_pageBlockCollage                                 TL_CLASS_ID = 145955919
	TL_CLASS_pageBlockSlideshow                               TL_CLASS_ID = 319588707
	TL_CLASS_pageBlockChannel                                 TL_CLASS_ID = -283684427
	TL_CLASS_pageBlockAudio                                   TL_CLASS_ID = 834148991
	TL_CLASS_pagePart                                         TL_CLASS_ID = -1908433218
	TL_CLASS_pageFull                                         TL_CLASS_ID = 1433323434
	TL_CLASS_phoneCallDiscardReasonMissed                     TL_CLASS_ID = -2048646399
	TL_CLASS_phoneCallDiscardReasonDisconnect                 TL_CLASS_ID = -527056480
	TL_CLASS_phoneCallDiscardReasonHangup                     TL_CLASS_ID = 1471006352
	TL_CLASS_phoneCallDiscardReasonBusy                       TL_CLASS_ID = -84416311
	TL_CLASS_dataJSON                                         TL_CLASS_ID = 2104790276
	TL_CLASS_labeledPrice                                     TL_CLASS_ID = -886477832
	TL_CLASS_invoice                                          TL_CLASS_ID = -1022713000
	TL_CLASS_paymentCharge                                    TL_CLASS_ID = -368917890
	TL_CLASS_postAddress                                      TL_CLASS_ID = 512535275
	TL_CLASS_paymentRequestedInfo                             TL_CLASS_ID = -1868808300
	TL_CLASS_paymentSavedCredentialsCard                      TL_CLASS_ID = -842892769
	TL_CLASS_webDocument                                      TL_CLASS_ID = -971322408
	TL_CLASS_inputWebDocument                                 TL_CLASS_ID = -1678949555
	TL_CLASS_inputWebFileLocation                             TL_CLASS_ID = -1036396922
	TL_CLASS_upload_webFile                                   TL_CLASS_ID = 568808380
	TL_CLASS_payments_paymentForm                             TL_CLASS_ID = 1062645411
	TL_CLASS_payments_validatedRequestedInfo                  TL_CLASS_ID = -784000893
	TL_CLASS_payments_paymentResult                           TL_CLASS_ID = 1314881805
	TL_CLASS_payments_paymentVerficationNeeded                TL_CLASS_ID = 1800845601
	TL_CLASS_payments_paymentReceipt                          TL_CLASS_ID = 1342771681
	TL_CLASS_payments_savedInfo                               TL_CLASS_ID = -74456004
	TL_CLASS_inputPaymentCredentialsSaved                     TL_CLASS_ID = -1056001329
	TL_CLASS_inputPaymentCredentials                          TL_CLASS_ID = 873977640
	TL_CLASS_inputPaymentCredentialsApplePay                  TL_CLASS_ID = 178373535
	TL_CLASS_inputPaymentCredentialsAndroidPay                TL_CLASS_ID = 2035705766
	TL_CLASS_account_tmpPassword                              TL_CLASS_ID = -614138572
	TL_CLASS_shippingOption                                   TL_CLASS_ID = -1239335713
	TL_CLASS_inputStickerSetItem                              TL_CLASS_ID = -6249322
	TL_CLASS_inputPhoneCall                                   TL_CLASS_ID = 506920429
	TL_CLASS_phoneCallEmpty                                   TL_CLASS_ID = 1399245077
	TL_CLASS_phoneCallWaiting                                 TL_CLASS_ID = 462375633
	TL_CLASS_phoneCallRequested                               TL_CLASS_ID = -2089411356
	TL_CLASS_phoneCallAccepted                                TL_CLASS_ID = 1828732223
	TL_CLASS_phoneCall                                        TL_CLASS_ID = -1660057
	TL_CLASS_phoneCallDiscarded                               TL_CLASS_ID = 1355435489
	TL_CLASS_phoneConnection                                  TL_CLASS_ID = -1655957568
	TL_CLASS_phoneCallProtocol                                TL_CLASS_ID = -1564789301
	TL_CLASS_phone_phoneCall                                  TL_CLASS_ID = -326966976
	TL_CLASS_upload_cdnFileReuploadNeeded                     TL_CLASS_ID = -290921362
	TL_CLASS_upload_cdnFile                                   TL_CLASS_ID = -1449145777
	TL_CLASS_cdnPublicKey                                     TL_CLASS_ID = -914167110
	TL_CLASS_cdnConfig                                        TL_CLASS_ID = 1462101002
	TL_CLASS_langPackString                                   TL_CLASS_ID = -892239370
	TL_CLASS_langPackStringPluralized                         TL_CLASS_ID = 1816636575
	TL_CLASS_langPackStringDeleted                            TL_CLASS_ID = 695856818
	TL_CLASS_langPackDifference                               TL_CLASS_ID = -209337866
	TL_CLASS_langPackLanguage                                 TL_CLASS_ID = 292985073
	TL_CLASS_channelAdminRights                               TL_CLASS_ID = 1568467877
	TL_CLASS_channelBannedRights                              TL_CLASS_ID = 1489977929
	TL_CLASS_channelAdminLogEventActionChangeTitle            TL_CLASS_ID = -421545947
	TL_CLASS_channelAdminLogEventActionChangeAbout            TL_CLASS_ID = 1427671598
	TL_CLASS_channelAdminLogEventActionChangeUsername         TL_CLASS_ID = 1783299128
	TL_CLASS_channelAdminLogEventActionChangePhoto            TL_CLASS_ID = -1204857405
	TL_CLASS_channelAdminLogEventActionToggleInvites          TL_CLASS_ID = 460916654
	TL_CLASS_channelAdminLogEventActionToggleSignatures       TL_CLASS_ID = 648939889
	TL_CLASS_channelAdminLogEventActionUpdatePinned           TL_CLASS_ID = -370660328
	TL_CLASS_channelAdminLogEventActionEditMessage            TL_CLASS_ID = 1889215493
	TL_CLASS_channelAdminLogEventActionDeleteMessage          TL_CLASS_ID = 1121994683
	TL_CLASS_channelAdminLogEventActionParticipantJoin        TL_CLASS_ID = 405815507
	TL_CLASS_channelAdminLogEventActionParticipantLeave       TL_CLASS_ID = -124291086
	TL_CLASS_channelAdminLogEventActionParticipantInvite      TL_CLASS_ID = -484690728
	TL_CLASS_channelAdminLogEventActionParticipantToggleBan   TL_CLASS_ID = -422036098
	TL_CLASS_channelAdminLogEventActionParticipantToggleAdmin TL_CLASS_ID = -714643696
	TL_CLASS_channelAdminLogEventActionChangeStickerSet       TL_CLASS_ID = -1312568665
	TL_CLASS_channelAdminLogEventActionTogglePreHistoryHidden TL_CLASS_ID = 1599903217
	TL_CLASS_channelAdminLogEvent                             TL_CLASS_ID = 995769920
	TL_CLASS_channels_adminLogResults                         TL_CLASS_ID = -309659827
	TL_CLASS_channelAdminLogEventsFilter                      TL_CLASS_ID = -368018716
	TL_CLASS_popularContact                                   TL_CLASS_ID = 1558266229
	TL_CLASS_cdnFileHash                                      TL_CLASS_ID = 2012136335
	TL_CLASS_messages_favedStickersNotModified                TL_CLASS_ID = -1634752813
	TL_CLASS_messages_favedStickers                           TL_CLASS_ID = -209768682
	TL_CLASS_recentMeUrlUnknown                               TL_CLASS_ID = 1189204285
	TL_CLASS_recentMeUrlUser                                  TL_CLASS_ID = -1917045962
	TL_CLASS_recentMeUrlChat                                  TL_CLASS_ID = -1608834311
	TL_CLASS_recentMeUrlChatInvite                            TL_CLASS_ID = -347535331
	TL_CLASS_recentMeUrlStickerSet                            TL_CLASS_ID = -1140172836
	TL_CLASS_help_recentMeUrls                                TL_CLASS_ID = 235081943
	TL_CLASS_inputSingleMedia                                 TL_CLASS_ID = 1588230153
	TL_CLASS_invokeAfterMsg                                   TL_CLASS_ID = -878758099
	TL_CLASS_invokeAfterMsgs                                  TL_CLASS_ID = 1036301552
	TL_CLASS_initConnection                                   TL_CLASS_ID = -951575130
	TL_CLASS_invokeWithLayer                                  TL_CLASS_ID = -627372787
	TL_CLASS_invokeWithoutUpdates                             TL_CLASS_ID = -1080796745
	TL_CLASS_auth_checkPhone                                  TL_CLASS_ID = 1877286395
	TL_CLASS_auth_sendCode                                    TL_CLASS_ID = -2035355412
	TL_CLASS_auth_signUp                                      TL_CLASS_ID = 453408308
	TL_CLASS_auth_signIn                                      TL_CLASS_ID = -1126886015
	TL_CLASS_auth_logOut                                      TL_CLASS_ID = 1461180992
	TL_CLASS_auth_resetAuthorizations                         TL_CLASS_ID = -1616179942
	TL_CLASS_auth_sendInvites                                 TL_CLASS_ID = 1998331287
	TL_CLASS_auth_exportAuthorization                         TL_CLASS_ID = -440401971
	TL_CLASS_auth_importAuthorization                         TL_CLASS_ID = -470837741
	TL_CLASS_auth_bindTempAuthKey                             TL_CLASS_ID = -841733627
	TL_CLASS_auth_importBotAuthorization                      TL_CLASS_ID = 1738800940
	TL_CLASS_auth_checkPassword                               TL_CLASS_ID = 174260510
	TL_CLASS_auth_requestPasswordRecovery                     TL_CLASS_ID = -661144474
	TL_CLASS_auth_recoverPassword                             TL_CLASS_ID = 1319464594
	TL_CLASS_auth_resendCode                                  TL_CLASS_ID = 1056025023
	TL_CLASS_auth_cancelCode                                  TL_CLASS_ID = 520357240
	TL_CLASS_auth_dropTempAuthKeys                            TL_CLASS_ID = -1907842680
	TL_CLASS_account_registerDevice                           TL_CLASS_ID = 1669245048
	TL_CLASS_account_unregisterDevice                         TL_CLASS_ID = 1707432768
	TL_CLASS_account_updateNotifySettings                     TL_CLASS_ID = -2067899501
	TL_CLASS_account_getNotifySettings                        TL_CLASS_ID = 313765169
	TL_CLASS_account_resetNotifySettings                      TL_CLASS_ID = -612493497
	TL_CLASS_account_updateProfile                            TL_CLASS_ID = 2018596725
	TL_CLASS_account_updateStatus                             TL_CLASS_ID = 1713919532
	TL_CLASS_account_getWallPapers                            TL_CLASS_ID = -1068696894
	TL_CLASS_account_reportPeer                               TL_CLASS_ID = -1374118561
	TL_CLASS_account_checkUsername                            TL_CLASS_ID = 655677548
	TL_CLASS_account_updateUsername                           TL_CLASS_ID = 1040964988
	TL_CLASS_account_getPrivacy                               TL_CLASS_ID = -623130288
	TL_CLASS_account_setPrivacy                               TL_CLASS_ID = -906486552
	TL_CLASS_account_deleteAccount                            TL_CLASS_ID = 1099779595
	TL_CLASS_account_getAccountTTL                            TL_CLASS_ID = 150761757
	TL_CLASS_account_setAccountTTL                            TL_CLASS_ID = 608323678
	TL_CLASS_account_sendChangePhoneCode                      TL_CLASS_ID = 149257707
	TL_CLASS_account_changePhone                              TL_CLASS_ID = 1891839707
	TL_CLASS_account_updateDeviceLocked                       TL_CLASS_ID = 954152242
	TL_CLASS_account_getAuthorizations                        TL_CLASS_ID = -484392616
	TL_CLASS_account_resetAuthorization                       TL_CLASS_ID = -545786948
	TL_CLASS_account_getPassword                              TL_CLASS_ID = 1418342645
	TL_CLASS_account_getPasswordSettings                      TL_CLASS_ID = -1131605573
	TL_CLASS_account_updatePasswordSettings                   TL_CLASS_ID = -92517498
	TL_CLASS_account_sendConfirmPhoneCode                     TL_CLASS_ID = 353818557
	TL_CLASS_account_confirmPhone                             TL_CLASS_ID = 1596029123
	TL_CLASS_account_getTmpPassword                           TL_CLASS_ID = 1250046590
	TL_CLASS_users_getUsers                                   TL_CLASS_ID = 227648840
	TL_CLASS_users_getFullUser                                TL_CLASS_ID = -902781519
	TL_CLASS_contacts_getStatuses                             TL_CLASS_ID = -995929106
	TL_CLASS_contacts_getContacts                             TL_CLASS_ID = -1071414113
	TL_CLASS_contacts_importContacts                          TL_CLASS_ID = 746589157
	TL_CLASS_contacts_deleteContact                           TL_CLASS_ID = -1902823612
	TL_CLASS_contacts_deleteContacts                          TL_CLASS_ID = 1504393374
	TL_CLASS_contacts_block                                   TL_CLASS_ID = 858475004
	TL_CLASS_contacts_unblock                                 TL_CLASS_ID = -448724803
	TL_CLASS_contacts_getBlocked                              TL_CLASS_ID = -176409329
	TL_CLASS_contacts_exportCard                              TL_CLASS_ID = -2065352905
	TL_CLASS_contacts_importCard                              TL_CLASS_ID = 1340184318
	TL_CLASS_contacts_search                                  TL_CLASS_ID = 301470424
	TL_CLASS_contacts_resolveUsername                         TL_CLASS_ID = -113456221
	TL_CLASS_contacts_getTopPeers                             TL_CLASS_ID = -728224331
	TL_CLASS_contacts_resetTopPeerRating                      TL_CLASS_ID = 451113900
	TL_CLASS_contacts_resetSaved                              TL_CLASS_ID = -2020263951
	TL_CLASS_messages_getMessages                             TL_CLASS_ID = 1109588596
	TL_CLASS_messages_getDialogs                              TL_CLASS_ID = 421243333
	TL_CLASS_messages_getHistory                              TL_CLASS_ID = -591691168
	TL_CLASS_messages_search                                  TL_CLASS_ID = 60726944
	TL_CLASS_messages_readHistory                             TL_CLASS_ID = 238054714
	TL_CLASS_messages_deleteHistory                           TL_CLASS_ID = 469850889
	TL_CLASS_messages_deleteMessages                          TL_CLASS_ID = -443640366
	TL_CLASS_messages_receivedMessages                        TL_CLASS_ID = 94983360
	TL_CLASS_messages_setTyping                               TL_CLASS_ID = -1551737264
	TL_CLASS_messages_sendMessage                             TL_CLASS_ID = -91733382
	TL_CLASS_messages_sendMedia                               TL_CLASS_ID = -923703407
	TL_CLASS_messages_forwardMessages                         TL_CLASS_ID = 1888354709
	TL_CLASS_messages_reportSpam                              TL_CLASS_ID = -820669733
	TL_CLASS_messages_hideReportSpam                          TL_CLASS_ID = -1460572005
	TL_CLASS_messages_getPeerSettings                         TL_CLASS_ID = 913498268
	TL_CLASS_messages_getChats                                TL_CLASS_ID = 1013621127
	TL_CLASS_messages_getFullChat                             TL_CLASS_ID = 998448230
	TL_CLASS_messages_editChatTitle                           TL_CLASS_ID = -599447467
	TL_CLASS_messages_editChatPhoto                           TL_CLASS_ID = -900957736
	TL_CLASS_messages_addChatUser                             TL_CLASS_ID = -106911223
	TL_CLASS_messages_deleteChatUser                          TL_CLASS_ID = -530505962
	TL_CLASS_messages_createChat                              TL_CLASS_ID = 164303470
	TL_CLASS_messages_forwardMessage                          TL_CLASS_ID = 865483769
	TL_CLASS_messages_getDhConfig                             TL_CLASS_ID = 651135312
	TL_CLASS_messages_requestEncryption                       TL_CLASS_ID = -162681021
	TL_CLASS_messages_acceptEncryption                        TL_CLASS_ID = 1035731989
	TL_CLASS_messages_discardEncryption                       TL_CLASS_ID = -304536635
	TL_CLASS_messages_setEncryptedTyping                      TL_CLASS_ID = 2031374829
	TL_CLASS_messages_readEncryptedHistory                    TL_CLASS_ID = 2135648522
	TL_CLASS_messages_sendEncrypted                           TL_CLASS_ID = -1451792525
	TL_CLASS_messages_sendEncryptedFile                       TL_CLASS_ID = -1701831834
	TL_CLASS_messages_sendEncryptedService                    TL_CLASS_ID = 852769188
	TL_CLASS_messages_receivedQueue                           TL_CLASS_ID = 1436924774
	TL_CLASS_messages_reportEncryptedSpam                     TL_CLASS_ID = 1259113487
	TL_CLASS_messages_readMessageContents                     TL_CLASS_ID = 916930423
	TL_CLASS_messages_getAllStickers                          TL_CLASS_ID = 479598769
	TL_CLASS_messages_getWebPagePreview                       TL_CLASS_ID = 623001124
	TL_CLASS_messages_exportChatInvite                        TL_CLASS_ID = 2106086025
	TL_CLASS_messages_checkChatInvite                         TL_CLASS_ID = 1051570619
	TL_CLASS_messages_importChatInvite                        TL_CLASS_ID = 1817183516
	TL_CLASS_messages_getStickerSet                           TL_CLASS_ID = 639215886
	TL_CLASS_messages_installStickerSet                       TL_CLASS_ID = -946871200
	TL_CLASS_messages_uninstallStickerSet                     TL_CLASS_ID = -110209570
	TL_CLASS_messages_startBot                                TL_CLASS_ID = -421563528
	TL_CLASS_messages_getMessagesViews                        TL_CLASS_ID = -993483427
	TL_CLASS_messages_toggleChatAdmins                        TL_CLASS_ID = -326379039
	TL_CLASS_messages_editChatAdmin                           TL_CLASS_ID = -1444503762
	TL_CLASS_messages_migrateChat                             TL_CLASS_ID = 363051235
	TL_CLASS_messages_searchGlobal                            TL_CLASS_ID = -1640190800
	TL_CLASS_messages_reorderStickerSets                      TL_CLASS_ID = 2016638777
	TL_CLASS_messages_getDocumentByHash                       TL_CLASS_ID = 864953444
	TL_CLASS_messages_searchGifs                              TL_CLASS_ID = -1080395925
	TL_CLASS_messages_getSavedGifs                            TL_CLASS_ID = -2084618926
	TL_CLASS_messages_saveGif                                 TL_CLASS_ID = 846868683
	TL_CLASS_messages_getInlineBotResults                     TL_CLASS_ID = 1364105629
	TL_CLASS_messages_setInlineBotResults                     TL_CLASS_ID = -346119674
	TL_CLASS_messages_sendInlineBotResult                     TL_CLASS_ID = -1318189314
	TL_CLASS_messages_getMessageEditData                      TL_CLASS_ID = -39416522
	TL_CLASS_messages_editMessage                             TL_CLASS_ID = 97630429
	TL_CLASS_messages_editInlineBotMessage                    TL_CLASS_ID = -1327463869
	TL_CLASS_messages_getBotCallbackAnswer                    TL_CLASS_ID = -2130010132
	TL_CLASS_messages_setBotCallbackAnswer                    TL_CLASS_ID = -712043766
	TL_CLASS_messages_getPeerDialogs                          TL_CLASS_ID = 764901049
	TL_CLASS_messages_saveDraft                               TL_CLASS_ID = -1137057461
	TL_CLASS_messages_getAllDrafts                            TL_CLASS_ID = 1782549861
	TL_CLASS_messages_getFeaturedStickers                     TL_CLASS_ID = 766298703
	TL_CLASS_messages_readFeaturedStickers                    TL_CLASS_ID = 1527873830
	TL_CLASS_messages_getRecentStickers                       TL_CLASS_ID = 1587647177
	TL_CLASS_messages_saveRecentSticker                       TL_CLASS_ID = 958863608
	TL_CLASS_messages_clearRecentStickers                     TL_CLASS_ID = -1986437075
	TL_CLASS_messages_getArchivedStickers                     TL_CLASS_ID = 1475442322
	TL_CLASS_messages_getMaskStickers                         TL_CLASS_ID = 1706608543
	TL_CLASS_messages_getAttachedStickers                     TL_CLASS_ID = -866424884
	TL_CLASS_messages_setGameScore                            TL_CLASS_ID = -1896289088
	TL_CLASS_messages_setInlineGameScore                      TL_CLASS_ID = 363700068
	TL_CLASS_messages_getGameHighScores                       TL_CLASS_ID = -400399203
	TL_CLASS_messages_getInlineGameHighScores                 TL_CLASS_ID = 258170395
	TL_CLASS_messages_getCommonChats                          TL_CLASS_ID = 218777796
	TL_CLASS_messages_getAllChats                             TL_CLASS_ID = -341307408
	TL_CLASS_messages_getWebPage                              TL_CLASS_ID = 852135825
	TL_CLASS_messages_toggleDialogPin                         TL_CLASS_ID = 847887978
	TL_CLASS_messages_reorderPinnedDialogs                    TL_CLASS_ID = -1784678844
	TL_CLASS_messages_getPinnedDialogs                        TL_CLASS_ID = -497756594
	TL_CLASS_messages_setBotShippingResults                   TL_CLASS_ID = -436833542
	TL_CLASS_messages_setBotPrecheckoutResults                TL_CLASS_ID = 163765653
	TL_CLASS_messages_uploadMedia                             TL_CLASS_ID = 1369162417
	TL_CLASS_messages_sendScreenshotNotification              TL_CLASS_ID = -914493408
	TL_CLASS_messages_getFavedStickers                        TL_CLASS_ID = 567151374
	TL_CLASS_messages_faveSticker                             TL_CLASS_ID = -1174420133
	TL_CLASS_messages_getUnreadMentions                       TL_CLASS_ID = 1180140658
	TL_CLASS_messages_readMentions                            TL_CLASS_ID = 251759059
	TL_CLASS_messages_getRecentLocations                      TL_CLASS_ID = 613691874
	TL_CLASS_messages_sendMultiMedia                          TL_CLASS_ID = 546656559
	TL_CLASS_messages_uploadEncryptedFile                     TL_CLASS_ID = 1347929239
	TL_CLASS_updates_getState                                 TL_CLASS_ID = -304838614
	TL_CLASS_updates_getDifference                            TL_CLASS_ID = 630429265
	TL_CLASS_updates_getChannelDifference                     TL_CLASS_ID = 51854712
	TL_CLASS_photos_updateProfilePhoto                        TL_CLASS_ID = -256159406
	TL_CLASS_photos_uploadProfilePhoto                        TL_CLASS_ID = 1328726168
	TL_CLASS_photos_deletePhotos                              TL_CLASS_ID = -2016444625
	TL_CLASS_photos_getUserPhotos                             TL_CLASS_ID = -1848823128
	TL_CLASS_upload_saveFilePart                              TL_CLASS_ID = -1291540959
	TL_CLASS_upload_getFile                                   TL_CLASS_ID = -475607115
	TL_CLASS_upload_saveBigFilePart                           TL_CLASS_ID = -562337987
	TL_CLASS_upload_getWebFile                                TL_CLASS_ID = 619086221
	TL_CLASS_upload_getCdnFile                                TL_CLASS_ID = 536919235
	TL_CLASS_upload_reuploadCdnFile                           TL_CLASS_ID = 452533257
	TL_CLASS_upload_getCdnFileHashes                          TL_CLASS_ID = -149567365
	TL_CLASS_help_getConfig                                   TL_CLASS_ID = -990308245
	TL_CLASS_help_getNearestDc                                TL_CLASS_ID = 531836966
	TL_CLASS_help_getAppUpdate                                TL_CLASS_ID = -1372724842
	TL_CLASS_help_saveAppLog                                  TL_CLASS_ID = 1862465352
	TL_CLASS_help_getInviteText                               TL_CLASS_ID = 1295590211
	TL_CLASS_help_getSupport                                  TL_CLASS_ID = -1663104819
	TL_CLASS_help_getAppChangelog                             TL_CLASS_ID = -1877938321
	TL_CLASS_help_getTermsOfService                           TL_CLASS_ID = 889286899
	TL_CLASS_help_setBotUpdatesStatus                         TL_CLASS_ID = -333262899
	TL_CLASS_help_getCdnConfig                                TL_CLASS_ID = 1375900482
	TL_CLASS_help_getRecentMeUrls                             TL_CLASS_ID = 1036054804
	TL_CLASS_channels_readHistory                             TL_CLASS_ID = -871347913
	TL_CLASS_channels_deleteMessages                          TL_CLASS_ID = -2067661490
	TL_CLASS_channels_deleteUserHistory                       TL_CLASS_ID = -787622117
	TL_CLASS_channels_reportSpam                              TL_CLASS_ID = -32999408
	TL_CLASS_channels_getMessages                             TL_CLASS_ID = -1814580409
	TL_CLASS_channels_getParticipants                         TL_CLASS_ID = 306054633
	TL_CLASS_channels_getParticipant                          TL_CLASS_ID = 1416484774
	TL_CLASS_channels_getChannels                             TL_CLASS_ID = 176122811
	TL_CLASS_channels_getFullChannel                          TL_CLASS_ID = 141781513
	TL_CLASS_channels_createChannel                           TL_CLASS_ID = -192332417
	TL_CLASS_channels_editAbout                               TL_CLASS_ID = 333610782
	TL_CLASS_channels_editAdmin                               TL_CLASS_ID = 548962836
	TL_CLASS_channels_editTitle                               TL_CLASS_ID = 1450044624
	TL_CLASS_channels_editPhoto                               TL_CLASS_ID = -248621111
	TL_CLASS_channels_checkUsername                           TL_CLASS_ID = 283557164
	TL_CLASS_channels_updateUsername                          TL_CLASS_ID = 890549214
	TL_CLASS_channels_joinChannel                             TL_CLASS_ID = 615851205
	TL_CLASS_channels_leaveChannel                            TL_CLASS_ID = -130635115
	TL_CLASS_channels_inviteToChannel                         TL_CLASS_ID = 429865580
	TL_CLASS_channels_exportInvite                            TL_CLASS_ID = -950663035
	TL_CLASS_channels_deleteChannel                           TL_CLASS_ID = -1072619549
	TL_CLASS_channels_toggleInvites                           TL_CLASS_ID = 1231065863
	TL_CLASS_channels_exportMessageLink                       TL_CLASS_ID = -934882771
	TL_CLASS_channels_toggleSignatures                        TL_CLASS_ID = 527021574
	TL_CLASS_channels_updatePinnedMessage                     TL_CLASS_ID = -1490162350
	TL_CLASS_channels_getAdminedPublicChannels                TL_CLASS_ID = -1920105769
	TL_CLASS_channels_editBanned                              TL_CLASS_ID = -1076292147
	TL_CLASS_channels_getAdminLog                             TL_CLASS_ID = 870184064
	TL_CLASS_channels_setStickers                             TL_CLASS_ID = -359881479
	TL_CLASS_channels_readMessageContents                     TL_CLASS_ID = -357180360
	TL_CLASS_channels_deleteHistory                           TL_CLASS_ID = -1355375294
	TL_CLASS_channels_togglePreHistoryHidden                  TL_CLASS_ID = -356796084
	TL_CLASS_bots_sendCustomRequest                           TL_CLASS_ID = -1440257555
	TL_CLASS_bots_answerWebhookJSONQuery                      TL_CLASS_ID = -434028723
	TL_CLASS_payments_getPaymentForm                          TL_CLASS_ID = -1712285883
	TL_CLASS_payments_getPaymentReceipt                       TL_CLASS_ID = -1601001088
	TL_CLASS_payments_validateRequestedInfo                   TL_CLASS_ID = 1997180532
	TL_CLASS_payments_sendPaymentForm                         TL_CLASS_ID = 730364339
	TL_CLASS_payments_getSavedInfo                            TL_CLASS_ID = 578650699
	TL_CLASS_payments_clearSavedInfo                          TL_CLASS_ID = -667062079
	TL_CLASS_stickers_createStickerSet                        TL_CLASS_ID = -1680314774
	TL_CLASS_stickers_removeStickerFromSet                    TL_CLASS_ID = -143257775
	TL_CLASS_stickers_changeStickerPosition                   TL_CLASS_ID = -4795190
	TL_CLASS_stickers_addStickerToSet                         TL_CLASS_ID = -2041315650
	TL_CLASS_phone_getCallConfig                              TL_CLASS_ID = 1430593449
	TL_CLASS_phone_requestCall                                TL_CLASS_ID = 1536537556
	TL_CLASS_phone_acceptCall                                 TL_CLASS_ID = 1003664544
	TL_CLASS_phone_confirmCall                                TL_CLASS_ID = 788404002
	TL_CLASS_phone_receivedCall                               TL_CLASS_ID = 399855457
	TL_CLASS_phone_discardCall                                TL_CLASS_ID = 2027164582
	TL_CLASS_phone_setCallRating                              TL_CLASS_ID = 475228724
	TL_CLASS_phone_saveCallDebug                              TL_CLASS_ID = 662363518
	TL_CLASS_langpack_getLangPack                             TL_CLASS_ID = -1699363442
	TL_CLASS_langpack_getStrings                              TL_CLASS_ID = 773776152
	TL_CLASS_langpack_getDifference                           TL_CLASS_ID = 187583869
	TL_CLASS_langpack_getLanguages                            TL_CLASS_ID = -2146445955
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
