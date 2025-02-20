# sew

`sew` is a command-line interface tool designed to manage Steam Workshop items, blueprints, mods, scripts, and other related tasks for the game **Space Engineers**. This tool simplifies interactions with the Steam Workshop, making it easier for players to upload and manage their creations.

> **Disclaimer**: This README was generated by ChatGPT and may not contain 100% accurate information. Please verify details and make any necessary adjustments.

## Features

- **Persistent SteamCMD Setup**: Automatically sets up a persistent SteamCMD in `%appdata%/SpaceEngineers/.steamcmd` and uses it if found.
- **Upload and Update Workshop Items**: Supports uploading and updating blueprints and mods to the Steam Workshop.
- **Metadata Management**:
  - `workshop.vdf`: Auto-generated for Steam Workshop uploads (should not be altered by users).
  - `info.txt`: Contains item metadata.
    - **Line 1**: Workshop item title.
    - **Lines 2+**: Description formatted in [Steam markup](https://steamcommunity.com/comment/Recommendation/formattinghelp).
- **Download Workshop Items**: Downloads items to the SteamCMD folder.
- **Fix Metadata**: Ensures local metadata (`info.txt`) matches the correct Workshop ID.
- **Steam Client Management**:
  - **`vent-steam`**: Closes Steam using `taskkill` and restarts it via the Steam protocol handler.
  - **`login`**: Signs into SteamCMD (a required step before using `sew`).

> **Note:**
> - `sew` does not handle batch operations directly but can be used in scripts for automation.
> - Deleting or unlisting Workshop items is not currently supported (planned for future updates).
> - Blueprints uploaded using this tool have a five-minute verification period before they become publicly accessible. This is a security measure by Steam to prevent malicious uploads.

## Installation

To install `sew`, use the following command:

```bash
go install github.com/Merith-TK/se-workshop/cmd/sew@main
```

Make sure you have Go installed and properly set up on your system. You can download and install Go from the [official website](https://golang.org/dl/).

## Usage

### Logging into SteamCMD

Before uploading or managing Workshop items, you must log into SteamCMD:

```bash
sew login <username> [password] [steamauth]
```

- `<username>` is **required**.
- `[password]` and `[steamauth]` are **optional**; if not provided, SteamCMD will prompt for them interactively.

### Uploading a Blueprint

1. **Locate Your Blueprint Folder**
   Run the following command to find the blueprint folder:
   ```bash
   sew bp folder
   ```

2. **Prepare Your Files**
   - Ensure your blueprint files are in the correct location.
   - In the blueprint folder, create an `info.txt` file alongside your `bp.sbc` file with the following format:
     - **Line 1**: Workshop item title.
     - **Line 2 and beyond**: Workshop item description (supports [Steam Markup](https://steamcommunity.com/comment/Recommendation/formattinghelp)).

3. **Upload/Update the Blueprint**
   Open **Command Prompt** or **PowerShell** in your blueprint folder and run:
   ```bash
   sew upload <path to the bp.sbc or folder containing bp.sbc>
   ```
   If the upload or update is successful, the tool will print the URL to the blueprint.

### Quick Command Rundown

- **`login <username> [password] [steamauth]`**
  - Log into SteamCMD. This is required for uploading or downloading Workshop items.

- **(`bp`, `blueprint`, `schematic`)**
  - Manage blueprint-related tasks. 
  - `sew bp folder` is the only implemented subcommand, and returns the path to your blueprints folder, useful for a quick reference or use in scripts

- **Mod Commands (`mod`, `mods`)**
  - Manage mod-related tasks.
  Not yet Implemented

- **Script Commands (`script`, `scripts`, `scr`, `src`)**
  - Handle in-game scripts.
  Not yet Implemented

- **`cmd`**
  - Directly interact with SteamCMD.
  - `sew cmd <args>` is the same as `steamcmd <args>`,
  - Uses the steamcmd folder located at `%APPDATA%/SpaceEngineers/.steamcmd`

- **(`dl`, `download`) `<workshopid>`**
  - Download a Workshop item to the SteamCMD folder.
  Note, does not parse a workshop url like `https://steamcommunity.com/sharedfiles/filedetails/?id=3353622153`, it only accepts the `3353622153` portion of the URL. 

- **`get-id`**
  - Retrieve and print the Workshop URL for a given item.

- **`set-id` `<workshopID>`**
  - Modify local metadata (`workshop.vdf` and `bp.sbc`) to associate the file with its Workshop ID. This helps determine whether to upload a new item or update an existing one.

- **(`upload`, `update`) `[path to, or folder containing, bp.sbc]`**
  - Upload a new Workshop item or update an existing one.
  - Uses current directory if no argument is given.

- **`vent-steam`**
  - Forcefully close the Steam client using `taskkill` and restart it using the protocol handler.
  - Using SteamCMD causes your Steam Client session to go "offline", something about local authentication causes this issue, this command was made specifically to just quickly restart steam from the runbox (`win+r`, `sew vent-steam`)

## Configuration

The application uses the following environment variable to determine the Space Engineers data directory:

- **`APPDATA`**: This should be set to the AppData folder for your user profile.

## Structure

The application is structured into several packages:

- **`shared`**: Contains common functionalities, configurations, and structures used across the application.
- **`sebp`**: Handles commands related to blueprints.
- **`semod`**: Manages commands related to mods.
- **`sescr`**: Deals with commands for in-game scripts.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
