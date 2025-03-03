package usecase

func (u usecase) HandleWebhook(event interface{}) error {
	// switch event.(type) {
	// case *github.PullRequestEvent:
	// 	payload := event.(*github.PullRequestEvent)
	// 	fmt.Println(payload.GetAction())
	// 	fmt.Println(payload.GetNumber())
	// 	fmt.Println(payload.PullRequest.Assignee.GetLogin())
	// 	fmt.Println(payload.Repo.GetFullName())
	// 	fmt.Println(payload.PullRequest.Head.GetRef())
	// 	fmt.Println(payload.PullRequest.Base.GetRef())
	// 	break

	// case *github.Comment:

	// 	break
	// case *github.ForkEvent:
	// 	payload := event.(*github.ForkEvent)
	// 	fmt.Println(payload.GetForkee().Owner.GetEmail())
	// 	fmt.Println(payload.GetForkee().Owner.GetName())
	// 	break
	// }
	return nil
}
