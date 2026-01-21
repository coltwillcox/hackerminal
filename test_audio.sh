#!/bin/bash
# Test audio playback for hackerminal

echo "Testing audio file playback..."
echo ""

# Test if audio file exists
if [ ! -f "assets/beep.wav" ]; then
    echo "ERROR: assets/beep.wav not found!"
    exit 1
fi

echo "✓ Audio file exists: $(ls -lh assets/beep.wav | awk '{print $5}')"

# Check for audio players
echo ""
echo "Checking for audio players..."
players=("paplay" "aplay" "ffplay" "mpv" "mplayer" "cvlc" "afplay")
found_player=""

for player in "${players[@]}"; do
    if command -v "$player" &> /dev/null; then
        echo "✓ Found: $player"
        if [ -z "$found_player" ]; then
            found_player="$player"
        fi
    fi
done

if [ -z "$found_player" ]; then
    echo "WARNING: No audio player found. Will fall back to terminal bell."
    exit 0
fi

echo ""
echo "Testing playback with $found_player..."

case "$found_player" in
    paplay|aplay)
        $found_player assets/beep.wav &
        ;;
    ffplay)
        ffplay -nodisp -autoexit -loglevel quiet assets/beep.wav &
        ;;
    mpv)
        mpv --no-video --really-quiet assets/beep.wav &
        ;;
    mplayer)
        mplayer -really-quiet -novideo assets/beep.wav &
        ;;
    cvlc)
        cvlc --play-and-exit --quiet assets/beep.wav &
        ;;
    afplay)
        afplay assets/beep.wav &
        ;;
esac

pid=$!
echo "✓ Audio playback started (PID: $pid)"

# Wait for playback to finish
wait $pid 2>/dev/null

echo "✓ Audio playback test complete!"
echo ""
echo "If you heard a bell sound, the audio system is working correctly."
