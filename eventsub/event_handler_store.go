package eventsub

import (
	"github.com/Adeithe/go-twitch/eventsub/events"
)

type handlerStore struct {
	channelBanEventHandlers                []func(events.ChannelBanEvent)
	channelCharityDonateHandlers           []func(events.ChannelCharityDonateEvent)
	channelCharityStartHandlers            []func(events.ChannelCharityStartEvent)
	channelCharityProgressHandlers         []func(events.ChannelCharityProgressEvent)
	channelCharityStopHandlers             []func(events.ChannelCharityStopEvent)
	channelCheerHandlers                   []func(events.ChannelCheerEvent)
	channelFollowHandlers                  []func(events.ChannelFollowEvent)
	channelGoalBeginHandlers               []func(events.ChannelGoalBeginEvent)
	channelGoalEndHandlers                 []func(events.ChannelGoalEndEvent)
	channelGoalProgressHandlers            []func(events.ChannelGoalProgressEvent)
	channelGuestStarGuestUpdateHandlers    []func(events.ChannelGuestStarGuestUpdateEvent)
	channelGuestStarSessionBeginHandlers   []func(events.ChannelGuestStarSessionBeginEvent)
	channelGuestStarSessionEndHandlers     []func(events.ChannelGuestStarSessionEndEvent)
	channelGuestStarSettingsUpdateHandlers []func(events.ChannelGuestStarSettingsUpdateEvent)
	channelGuestStarSlotUpdateHandlers     []func(events.ChannelGuestStarSlotUpdateEvent)
	ChannelHypeTrainBeginHandlers          []func(events.ChannelHypeTrainBeginEvent)
	channelHypeTrainEndHandlers            []func(events.ChannelHypeTrainEndEvent)
	channelHypeTrainProgressHandlers       []func(events.ChannelHypeTrainProgressEvent)
	channelModeratorAddHandlers            []func(events.ChannelModPromotionEvent)
	channelModeratorRemoveHandlers         []func(events.ChannelModDemotionEvent)
	channelPointsRedemptionAddHandlers     []func(events.ChannelPointsRedemptionAddEvent)
	channelPointsRedemptionUpdateHandlers  []func(events.ChannelPointsRedemptionUpdateEvent)
	channelPointsRewardAddHandlers         []func(events.ChannelPointRewardsAddEvent)
	channelPointsRewardsRemoveHandlers     []func(events.ChannelPointRewardsRemoveEvent)
	channelPointsRewardsUpdateHandlers     []func(events.ChannelPointRewardsUpdateEvent)
	channelPollBeginHandlers               []func(events.ChannelPollBeginEvent)
	channelPollEndHandlers                 []func(events.ChannelPollEndEvent)
	channelPollProgressHandlers            []func(events.ChannelPollProgressEvent)
	channelPredictionBeginHandlers         []func(events.ChannelPredictionBeginEvent)
	channelPredictionEndHandlers           []func(events.ChannelPredictionEndEvent)
	channelPredictionLockHandlers          []func(events.ChannelPredictionLockEvent)
	channelPredictionProgressHandlers      []func(events.ChannelPredictionProgressEvent)
	channelRaidHandlers                    []func(events.ChannelRaidEvent)
	channelShieldModeBeginHandlers         []func(events.ChannelShieldModeBeginEvent)
	channelShieldModeEndHandlers           []func(events.ChannelShieldModeEndEvent)
	channelShoutoutCreateHandlers          []func(events.ChannelShoutoutCreateEvent)
	channelShoutoutReceiveHandlers         []func(events.ChannelShoutoutReceiveEvent)
	channelSubscribeHandlers               []func(events.ChannelSubscribeEvent)
	channelSubscriptionEndHandlers         []func(events.ChannelSubscriptionEndEvent)
	channelSubscriptionGiftHandlers        []func(events.ChannelSubscriptionGiftedEvent)
	channelSubscriptionMessageHandlers     []func(events.ChannelSubscriptionMessageEvent)
	channelUnbanHandlers                   []func(events.ChannelUnbanEvent)
	channelUpdateHandlers                  []func(events.ChannelUpdateEvent)
	streamOfflineHandlers                  []func(events.StreamOfflineEvent)
	streamOnlineHandlers                   []func(events.StreamOnlineEvent)
	userUpdateHandlers                     []func(events.UserUpdateEvent)
}

func (h *handlerStore) OnChannelBan(f func(events.ChannelBanEvent)) {
	h.channelBanEventHandlers = append(h.channelBanEventHandlers, f)
}

func (h *handlerStore) OnChannelUnban(f func(events.ChannelUnbanEvent)) {
	h.channelUnbanHandlers = append(h.channelUnbanHandlers, f)
}

func (h *handlerStore) OnChannelRaid(f func(events.ChannelRaidEvent)) {
	h.channelRaidHandlers = append(h.channelRaidHandlers, f)
}

func (h *handlerStore) OnChannelCheer(f func(events.ChannelCheerEvent)) {
	h.channelCheerHandlers = append(h.channelCheerHandlers, f)
}

func (h *handlerStore) OnChannelFollow(f func(events.ChannelFollowEvent)) {
	h.channelFollowHandlers = append(h.channelFollowHandlers, f)
}

func (h *handlerStore) OnChannelCharityDonate(f func(events.ChannelCharityDonateEvent)) {
	h.channelCharityDonateHandlers = append(h.channelCharityDonateHandlers, f)
}

func (h *handlerStore) OnChannelCharityProgress(f func(events.ChannelCharityProgressEvent)) {
	h.channelCharityProgressHandlers = append(h.channelCharityProgressHandlers, f)
}

func (h *handlerStore) OnChannelCharityStart(f func(events.ChannelCharityStartEvent)) {
	h.channelCharityStartHandlers = append(h.channelCharityStartHandlers, f)
}

func (h *handlerStore) OnChannelCharityStop(f func(events.ChannelCharityStopEvent)) {
	h.channelCharityStopHandlers = append(h.channelCharityStopHandlers, f)
}

func (h *handlerStore) OnChannelGoalBegin(f func(events.ChannelGoalBeginEvent)) {
	h.channelGoalBeginHandlers = append(h.channelGoalBeginHandlers, f)
}

func (h *handlerStore) OnChannelGoalEnd(f func(events.ChannelGoalEndEvent)) {
	h.channelGoalEndHandlers = append(h.channelGoalEndHandlers, f)
}

func (h *handlerStore) OnChannelGoalProgress(f func(events.ChannelGoalProgressEvent)) {
	h.channelGoalProgressHandlers = append(h.channelGoalProgressHandlers, f)
}

func (h *handlerStore) OnChannelGuestStarGuestGuestUpdate(f func(events.ChannelGuestStarGuestUpdateEvent)) {
	h.channelGuestStarGuestUpdateHandlers = append(h.channelGuestStarGuestUpdateHandlers, f)
}

func (h *handlerStore) OnChannelGuestStarSessionBegin(f func(events.ChannelGuestStarSessionBeginEvent)) {
	h.channelGuestStarSessionBeginHandlers = append(h.channelGuestStarSessionBeginHandlers, f)
}

func (h *handlerStore) OnChannelGuestStarSessionEnd(f func(events.ChannelGuestStarSessionEndEvent)) {
	h.channelGuestStarSessionEndHandlers = append(h.channelGuestStarSessionEndHandlers, f)
}

func (h *handlerStore) OnChannelGuestStarSettingsUpdate(f func(events.ChannelGuestStarSettingsUpdateEvent)) {
	h.channelGuestStarSettingsUpdateHandlers = append(h.channelGuestStarSettingsUpdateHandlers, f)
}

func (h *handlerStore) OnChannelGuestStarSlotUpdate(f func(events.ChannelGuestStarSlotUpdateEvent)) {
	h.channelGuestStarSlotUpdateHandlers = append(h.channelGuestStarSlotUpdateHandlers, f)
}

func (h *handlerStore) OnChannelHypeTrainBegin(f func(events.ChannelHypeTrainBeginEvent)) {
	h.ChannelHypeTrainBeginHandlers = append(h.ChannelHypeTrainBeginHandlers, f)
}

func (h *handlerStore) OnChannelHypeTrainEnd(f func(events.ChannelHypeTrainEndEvent)) {
	h.channelHypeTrainEndHandlers = append(h.channelHypeTrainEndHandlers, f)
}

func (h *handlerStore) OnChannelHypeTrainProgress(f func(events.ChannelHypeTrainProgressEvent)) {
	h.channelHypeTrainProgressHandlers = append(h.channelHypeTrainProgressHandlers, f)
}

func (h *handlerStore) OnChannelModeratorAdd(f func(events.ChannelModPromotionEvent)) {
	h.channelModeratorAddHandlers = append(h.channelModeratorAddHandlers, f)
}

func (h *handlerStore) OnChannelModeratorRemove(f func(events.ChannelModDemotionEvent)) {
	h.channelModeratorRemoveHandlers = append(h.channelModeratorRemoveHandlers, f)
}

func (h *handlerStore) OnChannelPointsRedemptionAdd(f func(events.ChannelPointsRedemptionAddEvent)) {
	h.channelPointsRedemptionAddHandlers = append(h.channelPointsRedemptionAddHandlers, f)
}

func (h *handlerStore) OnChannelPointsRedemptionUpdate(f func(events.ChannelPointsRedemptionUpdateEvent)) {
	h.channelPointsRedemptionUpdateHandlers = append(h.channelPointsRedemptionUpdateHandlers, f)
}

func (h *handlerStore) OnChannelPointsRewardsAdd(f func(events.ChannelPointRewardsAddEvent)) {
	h.channelPointsRewardAddHandlers = append(h.channelPointsRewardAddHandlers, f)
}

func (h *handlerStore) OnChannelPointsRewardsRemove(f func(events.ChannelPointRewardsRemoveEvent)) {
	h.channelPointsRewardsRemoveHandlers = append(h.channelPointsRewardsRemoveHandlers, f)
}

func (h *handlerStore) OnChannelPointsRewardsUpdate(f func(events.ChannelPointRewardsUpdateEvent)) {
	h.channelPointsRewardsUpdateHandlers = append(h.channelPointsRewardsUpdateHandlers, f)
}

func (h *handlerStore) OnChannelPollBegin(f func(events.ChannelPollBeginEvent)) {
	h.channelPollBeginHandlers = append(h.channelPollBeginHandlers, f)
}

func (h *handlerStore) OnChannelPollEnd(f func(events.ChannelPollEndEvent)) {
	h.channelPollEndHandlers = append(h.channelPollEndHandlers, f)
}

func (h *handlerStore) OnChannelPollProgress(f func(events.ChannelPollProgressEvent)) {
	h.channelPollProgressHandlers = append(h.channelPollProgressHandlers, f)
}

func (h *handlerStore) OnChannelPredictionBegin(f func(events.ChannelPredictionBeginEvent)) {
	h.channelPredictionBeginHandlers = append(h.channelPredictionBeginHandlers, f)
}

func (h *handlerStore) OnChannelPredictionEnd(f func(events.ChannelPredictionEndEvent)) {
	h.channelPredictionEndHandlers = append(h.channelPredictionEndHandlers, f)
}

func (h *handlerStore) OnChannelPredictionLock(f func(events.ChannelPredictionLockEvent)) {
	h.channelPredictionLockHandlers = append(h.channelPredictionLockHandlers, f)
}

func (h *handlerStore) OnChannelPredictionProgress(f func(events.ChannelPredictionProgressEvent)) {
	h.channelPredictionProgressHandlers = append(h.channelPredictionProgressHandlers, f)
}

func (h *handlerStore) OnChannelShieldModeBegin(f func(events.ChannelShieldModeBeginEvent)) {
	h.channelShieldModeBeginHandlers = append(h.channelShieldModeBeginHandlers, f)
}

func (h *handlerStore) OnChannelShieldModeEnd(f func(events.ChannelShieldModeEndEvent)) {
	h.channelShieldModeEndHandlers = append(h.channelShieldModeEndHandlers, f)
}

func (h *handlerStore) OnChannelShoutoutCreate(f func(events.ChannelShoutoutCreateEvent)) {
	h.channelShoutoutCreateHandlers = append(h.channelShoutoutCreateHandlers, f)
}

func (h *handlerStore) OnChannelShoutoutReceive(f func(events.ChannelShoutoutReceiveEvent)) {
	h.channelShoutoutReceiveHandlers = append(h.channelShoutoutReceiveHandlers, f)
}

func (h *handlerStore) OnChannelSubscribe(f func(events.ChannelSubscribeEvent)) {
	h.channelSubscribeHandlers = append(h.channelSubscribeHandlers, f)
}

func (h *handlerStore) OnChannelSubscriptionEnd(f func(events.ChannelSubscriptionEndEvent)) {
	h.channelSubscriptionEndHandlers = append(h.channelSubscriptionEndHandlers, f)
}

func (h *handlerStore) OnChannelSubscriptionGift(f func(events.ChannelSubscriptionGiftedEvent)) {
	h.channelSubscriptionGiftHandlers = append(h.channelSubscriptionGiftHandlers, f)
}

func (h *handlerStore) OnChannelSubscriptionMessage(f func(events.ChannelSubscriptionMessageEvent)) {
	h.channelSubscriptionMessageHandlers = append(h.channelSubscriptionMessageHandlers, f)
}

func (h *handlerStore) OnStreamOffline(f func(events.StreamOfflineEvent)) {
	h.streamOfflineHandlers = append(h.streamOfflineHandlers, f)
}

func (h *handlerStore) OnStreamOnline(f func(events.StreamOnlineEvent)) {
	h.streamOnlineHandlers = append(h.streamOnlineHandlers, f)
}

func (h *handlerStore) OnChannelUpdate(f func(events.ChannelUpdateEvent)) {
	h.channelUpdateHandlers = append(h.channelUpdateHandlers, f)
}

func (h *handlerStore) OnUserUpdate(f func(events.UserUpdateEvent)) {
	h.userUpdateHandlers = append(h.userUpdateHandlers, f)
}

func (h *handlerStore) sendChannelBan(e events.ChannelBanEvent) {
	for _, i := range h.channelBanEventHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelUnban(e events.ChannelUnbanEvent) {
	for _, i := range h.channelUnbanHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelRaid(e events.ChannelRaidEvent) {
	for _, i := range h.channelRaidHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelCheer(e events.ChannelCheerEvent) {
	for _, i := range h.channelCheerHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelFollow(e events.ChannelFollowEvent) {
	for _, i := range h.channelFollowHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelCharityDonate(e events.ChannelCharityDonateEvent) {
	for _, i := range h.channelCharityDonateHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelCharityProgress(e events.ChannelCharityProgressEvent) {
	for _, i := range h.channelCharityProgressHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelCharityStart(e events.ChannelCharityStartEvent) {
	for _, i := range h.channelCharityStartHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelCharityStop(e events.ChannelCharityStopEvent) {
	for _, i := range h.channelCharityStopHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelGoalBegin(e events.ChannelGoalBeginEvent) {
	for _, i := range h.channelGoalBeginHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelGoalEnd(e events.ChannelGoalEndEvent) {
	for _, i := range h.channelGoalEndHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelGoalProgress(e events.ChannelGoalProgressEvent) {
	for _, i := range h.channelGoalProgressHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelGuestStarGuestGuestUpdate(e events.ChannelGuestStarGuestUpdateEvent) {
	for _, i := range h.channelGuestStarGuestUpdateHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelGuestStarSessionBegin(e events.ChannelGuestStarSessionBeginEvent) {
	for _, i := range h.channelGuestStarSessionBeginHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelGuestStarSessionEnd(e events.ChannelGuestStarSessionEndEvent) {
	for _, i := range h.channelGuestStarSessionEndHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelGuestStarSettingsUpdate(e events.ChannelGuestStarSettingsUpdateEvent) {
	for _, i := range h.channelGuestStarSettingsUpdateHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelGuestStarSlotUpdate(e events.ChannelGuestStarSlotUpdateEvent) {
	for _, i := range h.channelGuestStarSlotUpdateHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelHypeTrainBegin(e events.ChannelHypeTrainBeginEvent) {
	for _, i := range h.ChannelHypeTrainBeginHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelHypeTrainEnd(e events.ChannelHypeTrainEndEvent) {
	for _, i := range h.channelHypeTrainEndHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelHypeTrainProgress(e events.ChannelHypeTrainProgressEvent) {
	for _, i := range h.channelHypeTrainProgressHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelModeratorAdd(e events.ChannelModPromotionEvent) {
	for _, i := range h.channelModeratorAddHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelModeratorRemove(e events.ChannelModDemotionEvent) {
	for _, i := range h.channelModeratorRemoveHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPointsRedemptionAdd(e events.ChannelPointsRedemptionAddEvent) {
	for _, i := range h.channelPointsRedemptionAddHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPointsRedemptionUpdate(e events.ChannelPointsRedemptionUpdateEvent) {
	for _, i := range h.channelPointsRedemptionUpdateHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPointsRewardsAdd(e events.ChannelPointRewardsAddEvent) {
	for _, i := range h.channelPointsRewardAddHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPointsRewardsRemove(e events.ChannelPointRewardsRemoveEvent) {
	for _, i := range h.channelPointsRewardsRemoveHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPointsRewardsUpdate(e events.ChannelPointRewardsUpdateEvent) {
	for _, i := range h.channelPointsRewardsUpdateHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPollBegin(e events.ChannelPollBeginEvent) {
	for _, i := range h.channelPollBeginHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPollEnd(e events.ChannelPollEndEvent) {
	for _, i := range h.channelPollEndHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPollProgress(e events.ChannelPollProgressEvent) {
	for _, i := range h.channelPollProgressHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPredictionBegin(e events.ChannelPredictionBeginEvent) {
	for _, i := range h.channelPredictionBeginHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPredictionEnd(e events.ChannelPredictionEndEvent) {
	for _, i := range h.channelPredictionEndHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPredictionLock(e events.ChannelPredictionLockEvent) {
	for _, i := range h.channelPredictionLockHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelPredictionProgress(e events.ChannelPredictionProgressEvent) {
	for _, i := range h.channelPredictionProgressHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelShieldModeBegin(e events.ChannelShieldModeBeginEvent) {
	for _, i := range h.channelShieldModeBeginHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelShieldModeEnd(e events.ChannelShieldModeEndEvent) {
	for _, i := range h.channelShieldModeEndHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelShoutoutCreate(e events.ChannelShoutoutCreateEvent) {
	for _, i := range h.channelShoutoutCreateHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelShoutoutReceive(e events.ChannelShoutoutReceiveEvent) {
	for _, i := range h.channelShoutoutReceiveHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelSubscribe(e events.ChannelSubscribeEvent) {
	for _, i := range h.channelSubscribeHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelSubscriptionEnd(e events.ChannelSubscriptionEndEvent) {
	for _, i := range h.channelSubscriptionEndHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelSubscriptionGift(e events.ChannelSubscriptionGiftedEvent) {
	for _, i := range h.channelSubscriptionGiftHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelSubscriptionMessage(e events.ChannelSubscriptionMessageEvent) {
	for _, i := range h.channelSubscriptionMessageHandlers {
		i(e)
	}
}

func (h *handlerStore) sendStreamOffline(e events.StreamOfflineEvent) {
	for _, i := range h.streamOfflineHandlers {
		i(e)
	}
}

func (h *handlerStore) sendStreamOnline(e events.StreamOnlineEvent) {
	for _, i := range h.streamOnlineHandlers {
		i(e)
	}
}

func (h *handlerStore) sendChannelUpdate(e events.ChannelUpdateEvent) {
	for _, i := range h.channelUpdateHandlers {
		i(e)
	}
}

func (h *handlerStore) sendUserUpdate(e events.UserUpdateEvent) {
	for _, i := range h.userUpdateHandlers {
		i(e)
	}
}
