# Graylog Labs

## Install Graylog

1. Install and Configure : https://go2docs.graylog.org/1380099/downloading_and_installing_graylog/installing_graylog.html

2. Test Log

```bash
# After ensuring that your Graylog, create a Raw/Plaintext Input by navigating to your Graylog port.
# After that you can send a plain text message to the Graylog Raw/Plaintext TCP input running on port 5555 using the following command:
echo -n '{"version": "1.1", "host": "example.org", "short_message": "Test GELF TCP log", "level": 5}\0' | nc 0.0.0.0 12201
```

## Insight


## References

- [Graylog](https://go2docs.graylog.org/current/home.htm)