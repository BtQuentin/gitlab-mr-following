# Gitlab MR Following

This project display Gitlab MR once a day in Microsoft Teams.

## Usage

This is a very simple Go script to use.

Options are:

`-token` Gitlab API Token (require)

`-webhook` Teams Webhook for output (require)

`-team` Your team name (If your team name appears in MR Name) (require)

`-project` Gitlab project ID (require)

In all those options, some are exclusive:

- In every case, you need to set `token` and `webhook`
- If you want to get MR for a project:
    - `project` must be given
- If you want to get MR for a team:
    - `team` must be given