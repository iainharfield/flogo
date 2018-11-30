#!/bin/bash

git config --global user.email iain.harfield@ntlworld.com
git config --global user.name iainharfield   

git add activity.go activity.json activity_test.go
git commit -m "fixes"
git push -u origin master

