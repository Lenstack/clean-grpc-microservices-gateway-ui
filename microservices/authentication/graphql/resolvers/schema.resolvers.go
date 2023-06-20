package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"fmt"

	"github.com/Lenstack/lensaas-server/microservices/authentication/graphql/generated"
)

// SignIn is the resolver for the SignIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input generated.SignInRequest) (*generated.SignInResponse, error) {
	panic(fmt.Errorf("not implemented: SignIn - SignIn"))
}

// SignInCallback is the resolver for the SignInCallback field.
func (r *mutationResolver) SignInCallback(ctx context.Context, input generated.SignInCallbackRequest) (*generated.SignInCallbackResponse, error) {
	panic(fmt.Errorf("not implemented: SignInCallback - SignInCallback"))
}

// SignUp is the resolver for the SignUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input generated.SignUpRequest) (*generated.SignUpResponse, error) {
	panic(fmt.Errorf("not implemented: SignUp - SignUp"))
}

// SignOut is the resolver for the SignOut field.
func (r *mutationResolver) SignOut(ctx context.Context) (*generated.SignOutResponse, error) {
	panic(fmt.Errorf("not implemented: SignOut - SignOut"))
}

// ForgotPassword is the resolver for the ForgotPassword field.
func (r *mutationResolver) ForgotPassword(ctx context.Context, input generated.ForgotPasswordRequest) (*generated.ForgotPasswordResponse, error) {
	panic(fmt.Errorf("not implemented: ForgotPassword - ForgotPassword"))
}

// ResetPassword is the resolver for the ResetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, input generated.ResetPasswordRequest) (*generated.ResetPasswordResponse, error) {
	panic(fmt.Errorf("not implemented: ResetPassword - ResetPassword"))
}

// SendVerificationEmail is the resolver for the SendVerificationEmail field.
func (r *mutationResolver) SendVerificationEmail(ctx context.Context, input generated.SendVerificationEmailRequest) (*generated.SendVerificationEmailResponse, error) {
	panic(fmt.Errorf("not implemented: SendVerificationEmail - SendVerificationEmail"))
}

// VerifyEmail is the resolver for the VerifyEmail field.
func (r *mutationResolver) VerifyEmail(ctx context.Context, input generated.VerifyEmailRequest) (*generated.VerifyEmailResponse, error) {
	panic(fmt.Errorf("not implemented: VerifyEmail - VerifyEmail"))
}

// MultiFactor is the resolver for the MultiFactor field.
func (r *mutationResolver) MultiFactor(ctx context.Context, input generated.MultiFactorRequest) (*generated.MultiFactorResponse, error) {
	panic(fmt.Errorf("not implemented: MultiFactor - MultiFactor"))
}

// MagicLinkSignIn is the resolver for the MagicLinkSignIn field.
func (r *mutationResolver) MagicLinkSignIn(ctx context.Context, input generated.MagicLinkSignInRequest) (*generated.MagicLinkSignInResponse, error) {
	panic(fmt.Errorf("not implemented: MagicLinkSignIn - MagicLinkSignIn"))
}

// AiAssistant is the resolver for the AiAssistant field.
func (r *mutationResolver) AiAssistant(ctx context.Context, input generated.AiAssistantRequest) (*generated.AiAssistantResponse, error) {
	panic(fmt.Errorf("not implemented: AiAssistant - AiAssistant"))
}

// CheckEmailAvailability is the resolver for the CheckEmailAvailability field.
func (r *mutationResolver) CheckEmailAvailability(ctx context.Context, input generated.CheckEmailAvailabilityRequest) (*generated.CheckEmailAvailabilityResponse, error) {
	panic(fmt.Errorf("not implemented: CheckEmailAvailability - CheckEmailAvailability"))
}

// RefreshToken is the resolver for the RefreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context) (*generated.RefreshTokenResponse, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - RefreshToken"))
}

// Me is the resolver for the Me field.
func (r *queryResolver) Me(ctx context.Context) (*generated.User, error) {
	panic(fmt.Errorf("not implemented: Me - Me"))
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context) ([]*generated.User, error) {
	panic(fmt.Errorf("not implemented: Users - Users"))
}

// Organizations is the resolver for the Organizations field.
func (r *queryResolver) Organizations(ctx context.Context) ([]*generated.Organization, error) {
	panic(fmt.Errorf("not implemented: Organizations - Organizations"))
}

// Members is the resolver for the Members field.
func (r *queryResolver) Members(ctx context.Context) ([]*generated.Member, error) {
	panic(fmt.Errorf("not implemented: Members - Members"))
}

// Sessions is the resolver for the Sessions field.
func (r *queryResolver) Sessions(ctx context.Context) ([]*generated.Session, error) {
	panic(fmt.Errorf("not implemented: Sessions - Sessions"))
}

// Roles is the resolver for the Roles field.
func (r *queryResolver) Roles(ctx context.Context) ([]*generated.Role, error) {
	panic(fmt.Errorf("not implemented: Roles - Roles"))
}

// Permissions is the resolver for the Permissions field.
func (r *queryResolver) Permissions(ctx context.Context) ([]*generated.Permission, error) {
	panic(fmt.Errorf("not implemented: Permissions - Permissions"))
}

// Plans is the resolver for the Plans field.
func (r *queryResolver) Plans(ctx context.Context) ([]*generated.Plan, error) {
	panic(fmt.Errorf("not implemented: Plans - Plans"))
}

// Settings is the resolver for the Settings field.
func (r *queryResolver) Settings(ctx context.Context) ([]*generated.Setting, error) {
	panic(fmt.Errorf("not implemented: Settings - Settings"))
}

// Logs is the resolver for the Logs field.
func (r *subscriptionResolver) Logs(ctx context.Context) (<-chan *generated.Log, error) {
	panic(fmt.Errorf("not implemented: Logs - Logs"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
