<!-- LTeX: language=en-US -->
# cowboy-gorl
## Setting up the project for development
TODO

## Structure
- `assets`: Holds game assets.
- `documentation`: A place for further documentation media.
- `build`: Destination for build output.
- `main.go`: Application entry point.
- `makefile`: Manages build tasks.
- `pkg`: Go packages.
    - `animation`: Animate values over time using keyframes.
    - `logging`: Log out information at different levels.
    - `settings`: Load setting from a file and provide a fallback.
    - `entities`: A place for the games entities.
    - `render`: A custom rendering loop.
    - `scenes`: The games stages, composing together entities.
    - `util`: Various smaller functions, that might be useful everywhere.
- `scripts`: Utility scripts for development.

## Detailed Documentation

Detailed Documentation on many different subjects can be found here:
- [Scripts](/documentation/scripts.md)
- [Game Structure: Scenes & Entities](/documentation/scenes-and-entities.md)
- [Creating Scenes & Entities](/documentation/creating-scenes-and-entities.md)
- [GUI](/documentation/gui.md)
- [GUI Styledefs](/documentation/gui-styledef.md)
- [Audio](/documentation/audio.md)

## Assets in use
A list of all external assets currently in use, and a short note on usage conditions / licensing.

- [Azakaela Music Pack](https://azakaela.itch.io/azas-free-music-pack-2-wild-west-of-another-world) Under CC BY 4.0, credit 'Azakaela'

## Resources
- Sounds
    - https://blipsounds.com/community-library/

## Inspiration
- Pinterest: https://www.pinterest.de/kjuliusg/cowboy-gorl/
- Games:
    - Darkwood
    - Rainword (Visuals)
    - Microcivilization
    - Rimworld

- Youtube
    - MonsterGarden (https://www.youtube.com/@monlenz)

- Books
    - The Dark Tower (King)
