# Contribution Checklist

Thanks for your interest in contributing code!

Follow this checklist for submitting a pull request (PR):

1. Confirm that any required contributor agreements for Bingusboard have been satisfied. Repository guidelines are documented at [https://github.com/SomethingGeneric/bingusboard](https://github.com/SomethingGeneric/bingusboard).
2. Your ticket is a Help Wanted GitHub issue for the project you're contributing to.
    - If not, follow the process [here](contributions-without-ticket.md) or open an issue at [https://github.com/SomethingGeneric/bingusboard/issues](https://github.com/SomethingGeneric/bingusboard/issues).
3. Your code is thoroughly tested, including appropriate unit tests, and manual testing.
4. If applicable, user interface strings are included in the localization file ([en.json](https://github.com/mattermost/focalboard/blob/main/webapp/i18n/en.json))
    - In the webapp folder, run `npm run i18n-extract` to generate the new/updated strings.
5. The PR is submitted against the `main` branch from your fork.
6. The PR title begins with the GitHub Ticket ID (e.g. `[GH-394]`) and the summary template is filled out.

Once submitted, the automated build process must pass in order for the PR to be accepted. Any errors or failures need to be addressed in order for the PR to be accepted. Next, the PR goes through [code review](code-review.md). To learn about the review process for each project, read the [CONTRIBUTING.md](../CONTRIBUTING.md) file in this repository.
