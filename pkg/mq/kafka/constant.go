package kafka

import "time"

const (
	defaultDialTimeout  = 500 * time.Millisecond
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
)

const (
	TopicCommentPublish      = "comment_publish"
	TopicCommentCacheRebuild = "comment_cache_rebuild"
	TopicCommentOperator     = "comment_operator"

	TopicRelationFollow       = "relation_follow"
	TopicRelationCacheRebuild = "relation_cache_rebuild"
	TopicRelationSyncCount    = "relation_sync_count"
	TopicRelationOperator     = "relation_operator"

	TopicOpusOperator = "opus_operator"
)
