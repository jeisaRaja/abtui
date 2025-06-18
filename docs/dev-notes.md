# 📘 abtui – Audiobook TUI Player

A terminal-based audiobook player built with:
- Go
- Bubbletea (TUI)
- MPV (via JSON IPC)
- FFmpeg (for metadata/chapter parsing)

---

##  Planned Features

### 🔹 MVP
- [ ] Pick audiobook directories
- [ ] Scan and list audiobooks
- [ ] Store playback state in a file
- [ ] Resume playback from saved state
- [ ] Pause/seek/play via TUI
- [ ] State auto-saved on pause/exit
- [ ] Chapter navigation (parse via FFmpeg)
- [ ] Playback speed control (live + persistent)
- [ ] Multi-file book support (episodes/parts)
- [ ] Resume on launch
- [ ] Last activity sorting (recently played first)

### 🔸 Nice to Have
- [ ] Bookmarks (timestamp + note)
- [ ] Progress bar in TUI
- [ ] Search/filter books
- [ ] Mark book as finished/archive

### 🔧 Advanced
- [ ] Config file support (`~/.config/abtui/config.json`)
- [ ] Keyboard mapping customization
- [ ] Background/tmux mode
- [ ] Export/import state file
- [ ] Scrobble log (Goodreads/local)

---

## 🧠 `state.json` Structure

```json
{
  "playback_speed": 1.25,
  "current_book": "Dune",
  "books": {
    "Dune": {
      "title": "Dune",
      "path": "/audiobooks/dune",
      "episodes": ["dune_part1.m4b", "dune_part2.m4b"],
      "last_played": {
        "episode": 1,
        "position": 1423.4
      },
      "bookmarks": [
        {"position": 432.2, "note": "Cool quote about fear"},
        {"position": 892.0, "note": "Mentat explanation"}
      ],
      "finished": false
    }
  }
}
```

## Notes
- abtui assume audiobook will be stored as `title/book.m4b`
- read config, read state, scan and add new book to state

