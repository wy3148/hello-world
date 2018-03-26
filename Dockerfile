FROM centos:7

# This should have been done when build on Jenkins or any other automation tool
# RUN wget https://dl.google.com/go/go1.10.linux-amd64.tar.gz
# RUN tar -C /usr/local -xzf go1.10.linux-amd64.tar.gz
# RUN export PATH=$PATH:/usr/local/go/bin
# git clone checkout source code, and go build to get the 'hello-world' binary application file
# In this dockerfile, I directly copy the hello-world binary which was built on linux os

COPY ./hello-world /opt/hello-world
RUN chmod a+x /opt/hello-world
CMD ["/opt/hello-world"]