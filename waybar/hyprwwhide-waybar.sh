
# Get instance ID safely
INSTANCE_ID=$(hyprctl instances 2>/dev/null | grep -oP 'instance \K[^:]+')

if [[ -z "$INSTANCE_ID" ]]; then
    echo "Error: Could not find Hyprland instance"
    exit 1
fi

SOCKET_PATH="$XDG_RUNTIME_DIR/hypr/$INSTANCE_ID/.socket2.sock"

if [[ ! -S "$SOCKET_PATH" ]]; then
    echo "Error: Socket not found at $SOCKET_PATH"
    exit 1
fi

echo "Monitoring workspace changes via $SOCKET_PATH"
socat -U - UNIX-CONNECT:"$SOCKET_PATH" | while read -r line; do
    case "$line" in
    *"workspace>>"*)
        WORKSPACE_ID=$(echo "$line" | cut -d'>' -f3)
        # command here
        pkill -RTMIN+7 waybar
        ;;
    *"focusedmon>>"*)
        MONITOR=$(echo "$line" | cut -d'>' -f3)
        #echo "Monitor focus changed to $MONITOR"
        ;;
    esac
done