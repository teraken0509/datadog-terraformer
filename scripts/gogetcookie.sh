#!/bin/bash

touch ~/.gitcookies
chmod 0600 ~/.gitcookies

git config --global http.cookiefile ~/.gitcookies

tr , \\t <<\__END__ >>~/.gitcookies
.googlesource.com,TRUE,/,TRUE,2147483647,o,git-kterada.0509sg.gmail.com=1/yl1mDTt9G55PKFEJ3qrvRnWfL_nPaZhPDH9XnXCBabCfXaFC4NhbAlZiYM9oviu1
__END__