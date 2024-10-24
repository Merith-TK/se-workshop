# Usage

The `sew` command-line interface (CLI) provides a set of commands for managing Space Engineers blueprints, mods, scripts, and related Steam Workshop items. Below are the available commands and their usages:

## Commands

- **Blueprint Management**
  - `sew bp <command>` or `sew blueprints <command>`:
    - Commands related to managing blueprints.
  
- **Mod Management**
  - `sew mod <command>` or `sew mods <command>`:
    - Commands related to managing mods.
  
- **Script Management**
  - `sew script <command>` or `sew scripts <command>` or `sew scr <command>`:
    - Commands related to managing scripts.
  
- **SteamCMD Commands**
  - `sew cmd <args>`:
    - Executes SteamCMD commands.

- **Download Item**
  - `sew download <workshop_item_id>` or `sew dl <workshop_item_id>`:
    - Downloads a workshop item using its ID.

- **Get Workshop ID**
  - `sew get-id <path>` or `sew getid <path>` or `sew get <path>` or `sew id <path>`:
    - Retrieves the workshop ID from a provided path.

- **Get VDF File**
  - `sew get-vdf <path>` or `sew getvdf <path>` or `sew vdf <path>`:
    - Retrieves the VDF (Valve Data Format) file for the specified path.

- **Set Workshop ID**
  - `sew set-id <new_id>` or `sew setid <new_id>` or `sew set <new_id>`:
    - Sets a new workshop ID.

- **Fix Contents**
  - `sew fix-contents`:
    - Fixes the contents of the specified workshop item.

- **Upload/Update Item**
  - `sew upload <path>` or `sew update <path>`:
    - Uploads or updates a blueprint, mod, or script to the Steam Workshop.

- **Login to Steam**
  - `sew login <username> <password>`:
    - Logs into Steam using the provided credentials.

- **Vent Steam**
  - `sew vent-steam`:
    - Executes a specific command related to venting Steam.

## Notes

- If no arguments are provided, the application will display a help message.
- If only one argument is provided, the current working directory will be appended to the arguments.

### Examples

- To upload a blueprint:

  ```bash
  sew upload "path/to/blueprint"
  ```

- To download a workshop item:

  ```bash
  sew download 123456789
  ```

- To get the workshop ID of an item:

  ```bash
  sew get-id "path/to/item"
  ```

- To fix contents of an item:

  ```bash
  sew fix-contents
  ```
