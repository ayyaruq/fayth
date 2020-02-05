// Code generated by go generate; DO NOT EDIT.
// This file was generated at 05 Feb 20 19:45 AEDT,
// using data from https://github.com/SapphireServer/FFXIVOpcodes/raw/master/opcodes.json
package opcodes

var byCode = []map[uint32]string{
	101:   "PingHandler",
	978:   "InitHandler",
	476:   "FinishLoadingHandler",
	111:   "CFCommenceHandler",
	113:   "CFRegisterDuty",
	114:   "CFRegisterRoulette",
	630:   "PlayTimeHandler",
	726:   "LogoutHandler",
	143:   "CancelLogout",
	120:   "CFDutyInfoHandler",
	174:   "SocialReqSendHandler",
	175:   "CreateCrossWorldLS",
	767:   "ChatHandler",
	502:   "SocialListHandler",
	228:   "SetSearchInfoHandler",
	589:   "ReqSearchInfoHandler",
	231:   "ReqExamineSearchCommentHandler",
	241:   "ReqRemovePlayerFromBlacklist",
	359:   "BlackListHandler",
	244:   "PlayerSearchHandler",
	211:   "LinkshellListHandler",
	258:   "MarketBoardRequestItemListingInfo",
	259:   "MarketBoardRequestItemListings",
	263:   "MarketBoardSearch",
	275:   "ReqExamineFcInfo",
	282:   "FcInfoReqHandler",
	291:   "FreeCompanyUpdateShortMessageHandler",
	300:   "ReqMarketWishList",
	297:   "ReqJoinNoviceNetwork",
	309:   "ReqCountdownInitiate",
	310:   "ReqCountdownCancel",
	736:   "ZoneLineHandler",
	498:   "ClientTrigger",
	375:   "DiscoveryHandler",
	316:   "PlaceFieldMarker",
	317:   "SkillHandler",
	164:   "GMCommand1",
	319:   "GMCommand2",
	320:   "AoESkillHandler",
	447:   "UpdatePositionHandler",
	667:   "InventoryModifyHandler",
	329:   "InventoryEquipRecommendedItems",
	331:   "ReqPlaceHousingItem",
	335:   "BuildPresetHandler",
	337:   "TalkEventHandler",
	338:   "EmoteEventHandler",
	357:   "WithinRangeEventHandler",
	340:   "OutOfRangeEventHandler",
	341:   "EnterTeriEventHandler",
	342:   "ShopEventHandler",
	346:   "ReturnEventHandler",
	347:   "TradeReturnEventHandler",
	363:   "LinkshellEventHandler",
	364:   "LinkshellEventHandler1",
	373:   "ReqEquipDisplayFlagsChange",
	61815: "LandRenameHandler",
	376:   "HousingUpdateHouseGreeting",
	377:   "HousingUpdateObjectPosition",
	379:   "SetSharedEstateSettings",
	384:   "UpdatePositionInstance",
	667:   "PerformNoteHandler",
}

var byName = []map[string]uint32{
	"PingHandler":                          101,
	"InitHandler":                          978,
	"FinishLoadingHandler":                 476,
	"CFCommenceHandler":                    111,
	"CFRegisterDuty":                       113,
	"CFRegisterRoulette":                   114,
	"PlayTimeHandler":                      630,
	"LogoutHandler":                        726,
	"CancelLogout":                         143,
	"CFDutyInfoHandler":                    120,
	"SocialReqSendHandler":                 174,
	"CreateCrossWorldLS":                   175,
	"ChatHandler":                          767,
	"SocialListHandler":                    502,
	"SetSearchInfoHandler":                 228,
	"ReqSearchInfoHandler":                 589,
	"ReqExamineSearchCommentHandler":       231,
	"ReqRemovePlayerFromBlacklist":         241,
	"BlackListHandler":                     359,
	"PlayerSearchHandler":                  244,
	"LinkshellListHandler":                 211,
	"MarketBoardRequestItemListingInfo":    258,
	"MarketBoardRequestItemListings":       259,
	"MarketBoardSearch":                    263,
	"ReqExamineFcInfo":                     275,
	"FcInfoReqHandler":                     282,
	"FreeCompanyUpdateShortMessageHandler": 291,
	"ReqMarketWishList":                    300,
	"ReqJoinNoviceNetwork":                 297,
	"ReqCountdownInitiate":                 309,
	"ReqCountdownCancel":                   310,
	"ZoneLineHandler":                      736,
	"ClientTrigger":                        498,
	"DiscoveryHandler":                     375,
	"PlaceFieldMarker":                     316,
	"SkillHandler":                         317,
	"GMCommand1":                           164,
	"GMCommand2":                           319,
	"AoESkillHandler":                      320,
	"UpdatePositionHandler":                447,
	"InventoryModifyHandler":               667,
	"InventoryEquipRecommendedItems":       329,
	"ReqPlaceHousingItem":                  331,
	"BuildPresetHandler":                   335,
	"TalkEventHandler":                     337,
	"EmoteEventHandler":                    338,
	"WithinRangeEventHandler":              357,
	"OutOfRangeEventHandler":               340,
	"EnterTeriEventHandler":                341,
	"ShopEventHandler":                     342,
	"ReturnEventHandler":                   346,
	"TradeReturnEventHandler":              347,
	"LinkshellEventHandler":                363,
	"LinkshellEventHandler1":               364,
	"ReqEquipDisplayFlagsChange":           373,
	"LandRenameHandler":                    61815,
	"HousingUpdateHouseGreeting":           376,
	"HousingUpdateObjectPosition":          377,
	"SetSharedEstateSettings":              379,
	"UpdatePositionInstance":               384,
	"PerformNoteHandler":                   667,
}
