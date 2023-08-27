package events

// Interface for all events available
type EventReceiver interface {
	OnChannelBan(func(ChannelBanEvent))
	OnChannelUnban(func(ChannelUnbanEvent))

	OnChannelRaid(func(ChannelRaidEvent))
	OnChannelCheer(func(ChannelCheerEvent))
	OnChannelFollow(func(ChannelFollowEvent))

	OnChannelCharityDonate(func(ChannelCharityDonateEvent))
	OnChannelCharityProgress(func(ChannelCharityProgressEvent))
	OnChannelCharityStart(func(ChannelCharityStartEvent))
	OnChannelCharityStop(func(ChannelCharityStopEvent))

	OnChannelGoalBegin(func(ChannelGoalBeginEvent))
	OnChannelGoalEnd(func(ChannelGoalEndEvent))
	OnChannelGoalProgress(func(ChannelGoalProgressEvent))

	OnChannelGuestStarGuestGuestUpdate(func(ChannelGuestStarGuestUpdateEvent))
	OnChannelGuestStarSessionBegin(func(ChannelGuestStarSessionBeginEvent))
	OnChannelGuestStarSessionEnd(func(ChannelGuestStarSessionEndEvent))
	OnChannelGuestStarSettingsUpdate(func(ChannelGuestStarSettingsUpdateEvent))
	OnChannelGuestStarSlotUpdate(func(ChannelGuestStarSlotUpdateEvent))

	OnChannelHypeTrainBegin(func(ChannelHypeTrainBeginEvent))
	OnChannelHypeTrainEnd(func(ChannelHypeTrainEndEvent))
	OnChannelHypeTrainProgress(func(ChannelHypeTrainProgressEvent))

	OnChannelModeratorAdd(func(ChannelModPromotionEvent))
	OnChannelModeratorRemove(func(ChannelModDemotionEvent))

	OnChannelPointsRedemptionAdd(func(ChannelPointsRedemptionAddEvent))
	OnChannelPointsRedemptionUpdate(func(ChannelPointsRedemptionUpdateEvent))
	OnChannelPointsRewardsAdd(func(ChannelPointRewardsAddEvent))
	OnChannelPointsRewardsRemove(func(ChannelPointRewardsRemoveEvent))
	OnChannelPointsRewardsUpdate(func(ChannelPointRewardsUpdateEvent))

	OnChannelPollBegin(func(ChannelPollBeginEvent))
	OnChannelPollEnd(func(ChannelPollEndEvent))
	OnChannelPollProgress(func(ChannelPollProgressEvent))

	OnChannelPredictionBegin(func(ChannelPredictionBeginEvent))
	OnChannelPredictionEnd(func(ChannelPredictionEndEvent))
	OnChannelPredictionLock(func(ChannelPredictionLockEvent))
	OnChannelPredictionProgress(func(ChannelPredictionProgressEvent))

	OnChannelShieldModeActivate(func(ChannelShieldModeBeginEvent))
	OnChannelShieldModeEnd(func(ChannelShieldModeEndEvent))

	OnChannelShoutoutCreate(func(ChannelShoutoutCreateEvent))
	OnChannelShoutoutReceive(func(ChannelShoutoutReceiveEvent))

	OnChannelSubscribe(func(ChannelSubscribeEvent))
	OnChannelSubscriptionEnd(func(ChannelSubscriptionEndEvent))
	OnChannelSubscriptionGift(func(ChannelSubscriptionGiftedEvent))
	OnChannelSubscriptionMessage(func(ChannelSubscriptionMessageEvent))

	OnStreamOffline(func(StreamOfflineEvent))
	OnStreamOnline(func(StreamOnlineEvent))
	OnChannelUpdate(func(ChannelUpdateEvent))
	OnUserUpdate(func(UserUpdateEvent))
}
