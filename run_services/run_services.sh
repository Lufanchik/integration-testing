cat /dev/null > ../../logs/all.logs
#user-service
gnome-terminal --tab --title=user-service --command "bash -c 'echo \"Run user-service...\";./user-service.run;bash'" > /dev/null
#apm-gateway
gnome-terminal  --tab --title=apm-gateway --command "bash -c 'echo \"Run apm-api-gateway...\";./apm-api-gateway.run;bash'" > /dev/null
#auth-service
gnome-terminal  --tab --title=auth-service --command "bash -c 'echo \"Run auth-service...\";./auth-service.run;bash'" > /dev/null
#calculator-service
gnome-terminal  --tab --title=calculator-service --command "bash -c 'echo \"Run calculator-service...\";./calculator-service.run;bash'" > /dev/null
#pass-service
gnome-terminal  --tab --title=pass-service --command "bash -c 'echo \"Run pass-service...\";./pass-service.run;bash'" > /dev/null
#processing-api-gateway
#gnome-terminal  --tab --title=processing-api-gateway --command "bash -c 'echo \"Run processing-api-gateway...\";./processing-api-gateway.run;bash'" > /dev/null
#resolve-service
gnome-terminal  --tab --title=resolve-rmq --command "bash -c 'echo \"Run resolve-service-rmq...\";./resolve-service-rmq.run;bash'" > /dev/null
gnome-terminal  --tab --title=resolve-http --command "bash -c 'echo \"Run resolve-service-http...\";./resolve-service-http.run;bash'" > /dev/null
#revise-service
gnome-terminal  --tab --title=revise-http --command "bash -c 'echo \"Run revise-service-http...\";./revise-service-http.run;bash'" > /dev/null
gnome-terminal  --tab --title=revise-rmq --command "bash -c 'echo \"Run revise-service-rmq...\";./revise-service-rmq.run;bash'" > /dev/null
gnome-terminal  --tab --title=revise-task-watcher --command "bash -c 'echo \"Run revise-service-task-watcher...\";./revise-service-task-watcher.run;bash'" > /dev/null
#tidpool-service
gnome-terminal  --tab --title=tidpool-service --command "bash -c 'echo \"Run tidpool-service...\";./tidpool-service.run;bash'" > /dev/null
