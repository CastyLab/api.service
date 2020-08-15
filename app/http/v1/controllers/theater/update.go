package theater

import (
	"github.com/CastyLab/api.server/app/components"
	"github.com/CastyLab/api.server/app/http/v1/requests"
	"github.com/CastyLab/api.server/grpc"
	"github.com/CastyLab/grpc.proto/proto"
	"github.com/MrJoshLab/go-respond"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

func Update(ctx *gin.Context)  {

	var (
		req = &requests.UpdateTheaterRequest{
			Description: ctx.PostForm("description"),
		}
		token = ctx.Request.Header.Get("Authorization")
	)

	privacyInt, err := strconv.Atoi(ctx.PostForm("privacy"))
	if err == nil {
		req.Privacy = proto.PRIVACY(privacyInt)
	}

	videoPlayerAccessInt, err := strconv.Atoi(ctx.PostForm("video_player_access"))
	if err == nil {
		req.VideoPlayerAccess = proto.VIDEO_PLAYER_ACCESS(videoPlayerAccessInt)
	}

	if err := validator.New().Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		ctx.JSON(respond.Default.ValidationErrors(errors))
		return
	}

	_, err = grpc.TheaterServiceClient.UpdateTheater(ctx, &proto.TheaterAuthRequest{
		Theater: &proto.Theater{
			Description: req.Description,
			Privacy: req.Privacy,
			VideoPlayerAccess: req.VideoPlayerAccess,
		},
		AuthRequest: &proto.AuthenticateRequest{
			Token: []byte(token),
		},
	})

	if code, result, ok := components.ParseGrpcErrorResponse(err); !ok {
		ctx.JSON(code, result)
		return
	}

	ctx.JSON(respond.Default.UpdateSucceeded())
	return
}
