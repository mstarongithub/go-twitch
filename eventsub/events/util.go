package events

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// Convert a map m as delivered by json.Unmarshal into an event associated with the event type t
func ConvertMapToEvent(m map[string]interface{}, t string) (interface{}, error) {
	// Only way I can think of right now. A giant switch statement
	// And users will have to do the same too. Oh well
	switch t {
	case CHANNEL_BAN_EVENT:
		var x ChannelBanEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_CHARITY_DONATE_EVENT:
		var x ChannelCharityDonateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_CHARITY_START_EVENT:
		var x ChannelCharityStartEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_CHARITY_PROGRESS_EVENT:
		var x ChannelCharityProgressEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_CHARITY_STOP_EVENT:
		var x ChannelCharityStopEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_CHEER_EVENT:
		var x ChannelCheerEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_FOLLOW_EVENT:
		var x ChannelFollowEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_GOAL_BEGIN_EVENT:
		var x ChannelGoalBeginEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_GOAL_END_EVENT:
		var x ChannelGoalEndEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_GOAL_PROGRESS_EVENT:
		var x ChannelGoalProgressEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_GUEST_STAR_GUEST_UPDATE_EVENT:
		var x ChannelGuestStarGuestUpdateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_GUEST_STAR_SESSION_BEGIN_EVENT:
		var x ChannelGuestStarSessionBeginEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_GUEST_STAR_SESSION_END_EVENT:
		var x ChannelGuestStarSessionEndEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_GUEST_STAR_SETTINGS_UPDATE_EVENT:
		var x ChannelGuestStarSettingsUpdateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_GUEST_STAR_SLOT_UPDATE_EVENT:
		var x ChannelGuestStarSlotUpdateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_HYPE_TRAIN_BEGIN_EVENT:
		var x ChannelHypeTrainBeginEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_HYPE_TRAIN_END_EVENT:
		var x ChannelHypeTrainEndEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_HYPE_TRAIN_PROGRESS_EVENT:
		var x ChannelHypeTrainProgressEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_MODERATOR_ADD_EVENT:
		var x ChannelModPromotionEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_MODERATOR_REMOVE_EVENT:
		var x ChannelModDemotionEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_POINTS_REDEMPTION_ADD_EVENT:
		var x ChannelPointsRedemptionAddEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_POINTS_REDEMPTION_UPDATE_EVENT:
		var x ChannelPointsRedemptionUpdateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_POINTS_REWARD_ADD_EVENT:
		var x ChannelPointRewardsAddEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_POINTS_REWARDS_REMOVE_EVENT:
		var x ChannelPointRewardsRemoveEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_POINTS_REWARDS_UPDATE_EVENT:
		var x ChannelPointRewardsUpdateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_POLL_BEGIN_EVENT:
		var x ChannelPollBeginEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_POLL_END_EVENT:
		var x ChannelPollEndEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_POLL_PROGRESS_EVENT:
		var x ChannelPollProgressEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_PREDICTION_BEGIN_EVENT:
		var x ChannelPredictionBeginEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_PREDICTION_END_EVENT:
		var x ChannelPredictionEndEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_PREDICTION_LOCK_EVENT:
		var x ChannelPredictionLockEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_PREDICTION_PROGRESS_EVENT:
		var x ChannelPredictionProgressEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_RAID_EVENT:
		var x ChannelRaidEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_SHIELD_MODE_BEGIN_EVENT:
		var x ChannelShieldModeBeginEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_SHIELD_MODE_END_EVENT:
		var x ChannelShieldModeEndEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_SHOUTOUT_CREATE_EVENT:
		var x ChannelShoutoutCreateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_SHOUTOUT_RECEIVE_EVENT:
		var x ChannelShoutoutReceiveEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_SUBSCRIBE_EVENT:
		var x ChannelSubscribeEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_SUBSCRIPTION_END_EVENT:
		var x ChannelSubscriptionEndEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_SUBSCRIPTION_GIFT_EVENT:
		var x ChannelSubscriptionGiftedEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_SUBSCRIPTION_MESSAGE_EVENT:
		var x ChannelSubscriptionMessageEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_UNBAN_EVENT:
		var x ChannelUnbanEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case CHANNEL_UPDATE_EVENT:
		var x ChannelUpdateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case STREAM_OFFLINE_EVENT:
		var x StreamOfflineEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case STREAM_ONLINE_EVENT:
		var x StreamOnlineEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	case USER_UPDATE_EVENT:
		var x UserUpdateEvent
		err := mapstructure.Decode(m, &x)
		return x, err
	default:
		return nil, errors.New("not a valid event type")
	}
}
