FROM logstash

ADD logstash.conf /config-dir/logstash.conf
ADD patterns /patterns/

CMD ["-f", "/config-dir/logstash.conf"]