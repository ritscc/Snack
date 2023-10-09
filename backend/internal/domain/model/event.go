package model

// This event is used to notify the server or client side
// when data changes are made.
const (
	UserCreated = iota
	UserUpdated
	UserDeleted
	UserOnline
	UserOffline

	UserGroupCreated
	UserGroupUpdated
	UserGroupDeleted

	MemberAdded
	MemberUpdated
	MemberRemoved

	ChannelCreated
	ChannelUpdated
	ChannelDeleted

	MessageChannelJoined
	MessageChannelLeaved

	VoiceChannelJoined
	VoiceChannelLeaved

	ChannelEventCreated
	ChannelEventUpdated
	ChannelEventDeleted
	ChannelEventStamped
	ChannelEventUnstamped

	MessageCreated
	MessageUpdated
	MessageDeleted
	MessageStamped
	MessageUnstamped
	MessagePinned
	MessageUnpinned

	StampCreated
	StampUpdated
	StampDeleted
)
