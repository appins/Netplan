while true; do
  sleep 1
  if cat KILL_NETPLAN_NOW; then
    killall encode
    killall go
    killall make
    exit
  fi
done
