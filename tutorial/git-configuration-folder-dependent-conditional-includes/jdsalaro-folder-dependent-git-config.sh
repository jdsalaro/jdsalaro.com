#!/bin/bash
#set -e

mkdir /tmp/jdsalaro;
cd /tmp/jdsalaro;

for tldir in {'gitlab','gitlab-university','github'};
do

read -d "" gitconfig <<END

#_____________________________________________________________________________________________
#                                                                                            #
#                                                                                            #
# jdsalaro's git-configuration-folder-dependent-conditional-includes                         #
# from http://jdsalaro.com/tutorial/git-configuration-folder-dependent-conditional-includes  #
#                                                                                            # 
#____________________________________________________________________________________________#

[user]
	name = $tldir
	email = $tldir.noreply@jdsalaro.com
END

mkdir -p $tldir/repo
echo -e "$gitconfig" > $tldir/gitconfig
git init -q -b main $tldir/repo

done

tree


