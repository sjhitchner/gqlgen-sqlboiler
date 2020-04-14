// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package helpers

import (
	"github.com/web-ridge/gqlgen-sqlboiler/examples/social-network/models"
	"github.com/web-ridge/gqlgen-sqlboiler/helper"
)

var CommentPreloadMap = map[string]helper.ColumnSetting{
	"commentLikes": helper.ColumnSetting{
		Name:        models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"commentLikes.comment": helper.ColumnSetting{
		Name:        models.CommentRels.CommentLikes + "." + models.CommentLikeRels.Comment,
		IDAvailable: false,
	},
	"commentLikes.user": helper.ColumnSetting{
		Name:        models.CommentRels.CommentLikes + "." + models.CommentLikeRels.User,
		IDAvailable: false,
	},
	"post": helper.ColumnSetting{
		Name:        models.CommentRels.Post,
		IDAvailable: true,
	},
	"post.comments": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Comments,
		IDAvailable: false,
	},
	"post.comments.commentLikes": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"post.comments.post": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"post.comments.user": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"post.images": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Images,
		IDAvailable: false,
	},
	"post.images.imageVariations": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Images + "." + models.ImageRels.ImageVariations,
		IDAvailable: false,
	},
	"post.images.post": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Images + "." + models.ImageRels.Post,
		IDAvailable: false,
	},
	"post.likes": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Likes,
		IDAvailable: false,
	},
	"post.likes.post": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Likes + "." + models.LikeRels.Post,
		IDAvailable: false,
	},
	"post.likes.user": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.Likes + "." + models.LikeRels.User,
		IDAvailable: false,
	},
	"post.user": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.User,
		IDAvailable: false,
	},
	"post.user.commentLikes": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.User + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"post.user.comments": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.User + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"post.user.friendships": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.User + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"post.user.likes": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.User + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"post.user.posts": helper.ColumnSetting{
		Name:        models.CommentRels.Post + "." + models.PostRels.User + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
	"user": helper.ColumnSetting{
		Name:        models.CommentRels.User,
		IDAvailable: true,
	},
	"user.commentLikes": helper.ColumnSetting{
		Name:        models.CommentRels.User + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"user.comments": helper.ColumnSetting{
		Name:        models.CommentRels.User + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"user.friendships": helper.ColumnSetting{
		Name:        models.CommentRels.User + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"user.likes": helper.ColumnSetting{
		Name:        models.CommentRels.User + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"user.posts": helper.ColumnSetting{
		Name:        models.CommentRels.User + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
}

var CommentLikePreloadMap = map[string]helper.ColumnSetting{
	"comment": helper.ColumnSetting{
		Name:        models.CommentLikeRels.Comment,
		IDAvailable: true,
	},
	"comment.commentLikes": helper.ColumnSetting{
		Name:        models.CommentLikeRels.Comment + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"comment.commentLikes.comment": helper.ColumnSetting{
		Name:        models.CommentLikeRels.Comment + "." + models.CommentRels.CommentLikes + "." + models.CommentLikeRels.Comment,
		IDAvailable: false,
	},
	"comment.commentLikes.user": helper.ColumnSetting{
		Name:        models.CommentLikeRels.Comment + "." + models.CommentRels.CommentLikes + "." + models.CommentLikeRels.User,
		IDAvailable: false,
	},
	"comment.post": helper.ColumnSetting{
		Name:        models.CommentLikeRels.Comment + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"comment.user": helper.ColumnSetting{
		Name:        models.CommentLikeRels.Comment + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"user": helper.ColumnSetting{
		Name:        models.CommentLikeRels.User,
		IDAvailable: true,
	},
	"user.commentLikes": helper.ColumnSetting{
		Name:        models.CommentLikeRels.User + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"user.comments": helper.ColumnSetting{
		Name:        models.CommentLikeRels.User + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"user.friendships": helper.ColumnSetting{
		Name:        models.CommentLikeRels.User + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"user.likes": helper.ColumnSetting{
		Name:        models.CommentLikeRels.User + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"user.posts": helper.ColumnSetting{
		Name:        models.CommentLikeRels.User + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
}

var CommentLikePayloadPreloadLevels = struct {
	CommentLike string
}{
	CommentLike: "commentLike",
}

var CommentPayloadPreloadLevels = struct {
	Comment string
}{
	Comment: "comment",
}

var FriendshipPreloadMap = map[string]helper.ColumnSetting{
	"users": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users,
		IDAvailable: false,
	},
	"users.commentLikes": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"users.commentLikes.comment": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.CommentLikes + "." + models.CommentLikeRels.Comment,
		IDAvailable: false,
	},
	"users.commentLikes.user": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.CommentLikes + "." + models.CommentLikeRels.User,
		IDAvailable: false,
	},
	"users.comments": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"users.comments.commentLikes": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Comments + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"users.comments.post": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Comments + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"users.comments.post.comments": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Comments + "." + models.CommentRels.Post + "." + models.PostRels.Comments,
		IDAvailable: false,
	},
	"users.comments.post.images": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Comments + "." + models.CommentRels.Post + "." + models.PostRels.Images,
		IDAvailable: false,
	},
	"users.comments.post.likes": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Comments + "." + models.CommentRels.Post + "." + models.PostRels.Likes,
		IDAvailable: false,
	},
	"users.comments.post.user": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Comments + "." + models.CommentRels.Post + "." + models.PostRels.User,
		IDAvailable: false,
	},
	"users.comments.user": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Comments + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"users.friendships": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"users.friendships.users": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Friendships + "." + models.FriendshipRels.Users,
		IDAvailable: false,
	},
	"users.likes": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"users.likes.post": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Likes + "." + models.LikeRels.Post,
		IDAvailable: false,
	},
	"users.likes.user": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Likes + "." + models.LikeRels.User,
		IDAvailable: false,
	},
	"users.posts": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
	"users.posts.comments": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Posts + "." + models.PostRels.Comments,
		IDAvailable: false,
	},
	"users.posts.images": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Posts + "." + models.PostRels.Images,
		IDAvailable: false,
	},
	"users.posts.likes": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Posts + "." + models.PostRels.Likes,
		IDAvailable: false,
	},
	"users.posts.user": helper.ColumnSetting{
		Name:        models.FriendshipRels.Users + "." + models.UserRels.Posts + "." + models.PostRels.User,
		IDAvailable: false,
	},
}

var FriendshipPayloadPreloadLevels = struct {
	Friendship string
}{
	Friendship: "friendship",
}

var ImagePreloadMap = map[string]helper.ColumnSetting{
	"imageVariations": helper.ColumnSetting{
		Name:        models.ImageRels.ImageVariations,
		IDAvailable: false,
	},
	"imageVariations.image": helper.ColumnSetting{
		Name:        models.ImageRels.ImageVariations + "." + models.ImageVariationRels.Image,
		IDAvailable: false,
	},
	"imageVariations.image.imageVariations": helper.ColumnSetting{
		Name:        models.ImageRels.ImageVariations + "." + models.ImageVariationRels.Image + "." + models.ImageRels.ImageVariations,
		IDAvailable: false,
	},
	"imageVariations.image.post": helper.ColumnSetting{
		Name:        models.ImageRels.ImageVariations + "." + models.ImageVariationRels.Image + "." + models.ImageRels.Post,
		IDAvailable: false,
	},
	"post": helper.ColumnSetting{
		Name:        models.ImageRels.Post,
		IDAvailable: true,
	},
	"post.comments": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Comments,
		IDAvailable: false,
	},
	"post.comments.commentLikes": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"post.comments.post": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"post.comments.user": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"post.images": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Images,
		IDAvailable: false,
	},
	"post.images.imageVariations": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Images + "." + models.ImageRels.ImageVariations,
		IDAvailable: false,
	},
	"post.images.post": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Images + "." + models.ImageRels.Post,
		IDAvailable: false,
	},
	"post.likes": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Likes,
		IDAvailable: false,
	},
	"post.likes.post": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Likes + "." + models.LikeRels.Post,
		IDAvailable: false,
	},
	"post.likes.user": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.Likes + "." + models.LikeRels.User,
		IDAvailable: false,
	},
	"post.user": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.User,
		IDAvailable: false,
	},
	"post.user.commentLikes": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.User + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"post.user.comments": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.User + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"post.user.friendships": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.User + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"post.user.likes": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.User + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"post.user.posts": helper.ColumnSetting{
		Name:        models.ImageRels.Post + "." + models.PostRels.User + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
}

var ImagePayloadPreloadLevels = struct {
	Image string
}{
	Image: "image",
}

var ImageVariationPreloadMap = map[string]helper.ColumnSetting{
	"image": helper.ColumnSetting{
		Name:        models.ImageVariationRels.Image,
		IDAvailable: true,
	},
	"image.imageVariations": helper.ColumnSetting{
		Name:        models.ImageVariationRels.Image + "." + models.ImageRels.ImageVariations,
		IDAvailable: false,
	},
	"image.imageVariations.image": helper.ColumnSetting{
		Name:        models.ImageVariationRels.Image + "." + models.ImageRels.ImageVariations + "." + models.ImageVariationRels.Image,
		IDAvailable: false,
	},
	"image.post": helper.ColumnSetting{
		Name:        models.ImageVariationRels.Image + "." + models.ImageRels.Post,
		IDAvailable: false,
	},
	"image.post.comments": helper.ColumnSetting{
		Name:        models.ImageVariationRels.Image + "." + models.ImageRels.Post + "." + models.PostRels.Comments,
		IDAvailable: false,
	},
	"image.post.images": helper.ColumnSetting{
		Name:        models.ImageVariationRels.Image + "." + models.ImageRels.Post + "." + models.PostRels.Images,
		IDAvailable: false,
	},
	"image.post.likes": helper.ColumnSetting{
		Name:        models.ImageVariationRels.Image + "." + models.ImageRels.Post + "." + models.PostRels.Likes,
		IDAvailable: false,
	},
	"image.post.user": helper.ColumnSetting{
		Name:        models.ImageVariationRels.Image + "." + models.ImageRels.Post + "." + models.PostRels.User,
		IDAvailable: false,
	},
}

var ImageVariationPayloadPreloadLevels = struct {
	ImageVariation string
}{
	ImageVariation: "imageVariation",
}

var LikePreloadMap = map[string]helper.ColumnSetting{
	"post": helper.ColumnSetting{
		Name:        models.LikeRels.Post,
		IDAvailable: true,
	},
	"post.comments": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Comments,
		IDAvailable: false,
	},
	"post.comments.commentLikes": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"post.comments.post": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"post.comments.user": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Comments + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"post.images": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Images,
		IDAvailable: false,
	},
	"post.images.imageVariations": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Images + "." + models.ImageRels.ImageVariations,
		IDAvailable: false,
	},
	"post.images.post": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Images + "." + models.ImageRels.Post,
		IDAvailable: false,
	},
	"post.likes": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Likes,
		IDAvailable: false,
	},
	"post.likes.post": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Likes + "." + models.LikeRels.Post,
		IDAvailable: false,
	},
	"post.likes.user": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.Likes + "." + models.LikeRels.User,
		IDAvailable: false,
	},
	"post.user": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.User,
		IDAvailable: false,
	},
	"post.user.commentLikes": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.User + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"post.user.comments": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.User + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"post.user.friendships": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.User + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"post.user.likes": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.User + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"post.user.posts": helper.ColumnSetting{
		Name:        models.LikeRels.Post + "." + models.PostRels.User + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
	"user": helper.ColumnSetting{
		Name:        models.LikeRels.User,
		IDAvailable: true,
	},
	"user.commentLikes": helper.ColumnSetting{
		Name:        models.LikeRels.User + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"user.comments": helper.ColumnSetting{
		Name:        models.LikeRels.User + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"user.comments.commentLikes": helper.ColumnSetting{
		Name:        models.LikeRels.User + "." + models.UserRels.Comments + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"user.comments.post": helper.ColumnSetting{
		Name:        models.LikeRels.User + "." + models.UserRels.Comments + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"user.comments.user": helper.ColumnSetting{
		Name:        models.LikeRels.User + "." + models.UserRels.Comments + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"user.friendships": helper.ColumnSetting{
		Name:        models.LikeRels.User + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"user.likes": helper.ColumnSetting{
		Name:        models.LikeRels.User + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"user.posts": helper.ColumnSetting{
		Name:        models.LikeRels.User + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
}

var LikePayloadPreloadLevels = struct {
	Like string
}{
	Like: "like",
}

var PostPreloadMap = map[string]helper.ColumnSetting{
	"comments": helper.ColumnSetting{
		Name:        models.PostRels.Comments,
		IDAvailable: false,
	},
	"comments.commentLikes": helper.ColumnSetting{
		Name:        models.PostRels.Comments + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"comments.post": helper.ColumnSetting{
		Name:        models.PostRels.Comments + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"comments.user": helper.ColumnSetting{
		Name:        models.PostRels.Comments + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"images": helper.ColumnSetting{
		Name:        models.PostRels.Images,
		IDAvailable: false,
	},
	"images.imageVariations": helper.ColumnSetting{
		Name:        models.PostRels.Images + "." + models.ImageRels.ImageVariations,
		IDAvailable: false,
	},
	"images.post": helper.ColumnSetting{
		Name:        models.PostRels.Images + "." + models.ImageRels.Post,
		IDAvailable: false,
	},
	"likes": helper.ColumnSetting{
		Name:        models.PostRels.Likes,
		IDAvailable: false,
	},
	"likes.post": helper.ColumnSetting{
		Name:        models.PostRels.Likes + "." + models.LikeRels.Post,
		IDAvailable: false,
	},
	"likes.user": helper.ColumnSetting{
		Name:        models.PostRels.Likes + "." + models.LikeRels.User,
		IDAvailable: false,
	},
	"user": helper.ColumnSetting{
		Name:        models.PostRels.User,
		IDAvailable: true,
	},
	"user.commentLikes": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"user.commentLikes.comment": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.CommentLikes + "." + models.CommentLikeRels.Comment,
		IDAvailable: false,
	},
	"user.commentLikes.user": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.CommentLikes + "." + models.CommentLikeRels.User,
		IDAvailable: false,
	},
	"user.comments": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"user.comments.commentLikes": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Comments + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"user.comments.post": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Comments + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"user.comments.user": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Comments + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"user.friendships": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"user.likes": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"user.likes.post": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Likes + "." + models.LikeRels.Post,
		IDAvailable: false,
	},
	"user.likes.user": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Likes + "." + models.LikeRels.User,
		IDAvailable: false,
	},
	"user.posts": helper.ColumnSetting{
		Name:        models.PostRels.User + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
}

var PostPayloadPreloadLevels = struct {
	Post string
}{
	Post: "post",
}

var UserPreloadMap = map[string]helper.ColumnSetting{
	"commentLikes": helper.ColumnSetting{
		Name:        models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"commentLikes.comment": helper.ColumnSetting{
		Name:        models.UserRels.CommentLikes + "." + models.CommentLikeRels.Comment,
		IDAvailable: false,
	},
	"commentLikes.user": helper.ColumnSetting{
		Name:        models.UserRels.CommentLikes + "." + models.CommentLikeRels.User,
		IDAvailable: false,
	},
	"comments": helper.ColumnSetting{
		Name:        models.UserRels.Comments,
		IDAvailable: false,
	},
	"comments.commentLikes": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.CommentLikes,
		IDAvailable: false,
	},
	"comments.commentLikes.comment": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.CommentLikes + "." + models.CommentLikeRels.Comment,
		IDAvailable: false,
	},
	"comments.commentLikes.user": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.CommentLikes + "." + models.CommentLikeRels.User,
		IDAvailable: false,
	},
	"comments.post": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.Post,
		IDAvailable: false,
	},
	"comments.post.comments": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.Post + "." + models.PostRels.Comments,
		IDAvailable: false,
	},
	"comments.post.images": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.Post + "." + models.PostRels.Images,
		IDAvailable: false,
	},
	"comments.post.likes": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.Post + "." + models.PostRels.Likes,
		IDAvailable: false,
	},
	"comments.post.user": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.Post + "." + models.PostRels.User,
		IDAvailable: false,
	},
	"comments.user": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.User,
		IDAvailable: false,
	},
	"comments.user.commentLikes": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.User + "." + models.UserRels.CommentLikes,
		IDAvailable: false,
	},
	"comments.user.comments": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.User + "." + models.UserRels.Comments,
		IDAvailable: false,
	},
	"comments.user.friendships": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.User + "." + models.UserRels.Friendships,
		IDAvailable: false,
	},
	"comments.user.likes": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.User + "." + models.UserRels.Likes,
		IDAvailable: false,
	},
	"comments.user.posts": helper.ColumnSetting{
		Name:        models.UserRels.Comments + "." + models.CommentRels.User + "." + models.UserRels.Posts,
		IDAvailable: false,
	},
	"friendships": helper.ColumnSetting{
		Name:        models.UserRels.Friendships,
		IDAvailable: false,
	},
	"friendships.users": helper.ColumnSetting{
		Name:        models.UserRels.Friendships + "." + models.FriendshipRels.Users,
		IDAvailable: false,
	},
	"likes": helper.ColumnSetting{
		Name:        models.UserRels.Likes,
		IDAvailable: false,
	},
	"likes.post": helper.ColumnSetting{
		Name:        models.UserRels.Likes + "." + models.LikeRels.Post,
		IDAvailable: false,
	},
	"likes.user": helper.ColumnSetting{
		Name:        models.UserRels.Likes + "." + models.LikeRels.User,
		IDAvailable: false,
	},
	"posts": helper.ColumnSetting{
		Name:        models.UserRels.Posts,
		IDAvailable: false,
	},
	"posts.comments": helper.ColumnSetting{
		Name:        models.UserRels.Posts + "." + models.PostRels.Comments,
		IDAvailable: false,
	},
	"posts.images": helper.ColumnSetting{
		Name:        models.UserRels.Posts + "." + models.PostRels.Images,
		IDAvailable: false,
	},
	"posts.likes": helper.ColumnSetting{
		Name:        models.UserRels.Posts + "." + models.PostRels.Likes,
		IDAvailable: false,
	},
	"posts.user": helper.ColumnSetting{
		Name:        models.UserRels.Posts + "." + models.PostRels.User,
		IDAvailable: false,
	},
}

var UserPayloadPreloadLevels = struct {
	User string
}{
	User: "user",
}
