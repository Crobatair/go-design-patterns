## Proxy design pattern

The Proxy Pattern is a structural design pattern that is used to provide a surrogate or placeholder object, which references an underlying object. The Proxy Pattern provides a way to control access to the underlying object, and it can be used to add additional functionality to the underlying object.

Example:
- We have a large number of images that need to be loaded from a remote server. We can use the Proxy pattern to load the image only when it is needed.
    - We could have an Image object that loads the image from the remote server.
    - We could have a Proxy object that references the Image object.
    - The Proxy object can load the image from the remote server only when it is needed.
    - The Proxy object can also cache the image, so that it does not need to load the image from the remote server every time it is needed.

### When to use the Proxy Pattern
- When you need to control access to an object.
- When you need to add additional functionality to an object.
- When you need to defer the loading of an object until it is needed.
- When you need to cache an object.
- When you need to provide a placeholder for an object.