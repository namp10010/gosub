# gosub
a quick golang implementation of google pubsub subscriber for testing with the [emulator](https://cloud.google.com/pubsub/docs/emulator).

it does one thing which is subscribing to a topic and receiving messages. ok, that's two.

this can be used to read message your application sends to the emulator. there is a python script for it but as I can play around with Go and do more things I make this one. 

### install

```bash
  go install https://github.com/namp10010/gosub@latest
```

### usage

```bash
  PUBSUB_EMULATOR_HOST=localhost:8085 gosub -p project -t topic -s testsub
```

### links

* [pubsub emulator](https://cloud.google.com/pubsub/docs/emulator)