package fate_go_sdk

import (
	pb "github.com/zm-dev/fate-go-sdk/pb"
	"time"
	"context"
	"google.golang.org/grpc/metadata"
)

type LoginChecker struct {
	atc            *AccessTokenClient
	lcc            pb.LoginCheckerClient
	timeout        time.Duration
	accessTokenKey string
}

func (lc *LoginChecker) Check(ticketID string) (*pb.LoginCheckRes, error) {
	var (
		t   string
		err error
	)
	if t, err = lc.atc.GetToken(); err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), lc.timeout)
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{lc.accessTokenKey: t}))
	return lc.lcc.Check(ctx, &pb.TicketID{Id: ticketID})
}

func (lc *LoginChecker) Logout(ticketID string) (err error) {
	var t string
	if t, err = lc.atc.GetToken(); err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), lc.timeout)
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{lc.accessTokenKey: t}))
	_, err = lc.lcc.Logout(ctx, &pb.TicketID{Id: ticketID})
	return err
}

func NewLoginChecker(atc *AccessTokenClient, lcc pb.LoginCheckerClient, timeout time.Duration, accessTokenKey string) *LoginChecker {
	return &LoginChecker{atc: atc, lcc: lcc, timeout: timeout, accessTokenKey: accessTokenKey}
}
