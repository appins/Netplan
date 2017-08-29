while true; do
  sleep 1
  if cat KILL_NETPLAN_NOW; then
    killall make
    killall go
    killall encode
    exit
  fi
done
