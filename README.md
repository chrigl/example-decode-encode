This is just a small example to demonstrate testing of Encode/Decode. This should be relatively independend of the programming language, but for this example I used Go.


Just a short summary: I think you should always test Encode/Decode in combination, and never standalone. This ensures, that you won't accidentally break `Decode` because of a change in `Encode` and vice versa. Please browse the commit history for more information.
