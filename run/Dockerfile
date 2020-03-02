FROM golang:1.14

ENV GOPRIVATE="lab.siroccotechnology.ru"
RUN export PATH=$PATH:$HOME/bin
RUN echo "[url \"git@lab.siroccotechnology.ru:\"]\n\tinsteadOf = https://lab.siroccotechnology.ru" >> /root/.gitconfig

RUN mkdir -p ~/bin
RUN PATH=$PATH:/root/bin cd ~/bin && curl -L https://raw.githubusercontent.com/flant/multiwerf/master/get.sh | bash && . $(multiwerf use 1.0 stable --as-file)