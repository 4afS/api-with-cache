for i in `seq 100`
do
curl -o /dev/null http://localhost:8080/timeline/user1 &
sleep 0.001
done