# Statee internal APIs

The 'core' part being entirely responsible for fetching hardware information.
Designed to be easy and simple, yet verbose and scalable.

Can be used as an independent module.



## Special values

- **`-127` - unknown number**; returned when value couldn't be fetched or parsed properly due to an internal error; sometimes such behaviour is expected



## Structure

Each directory is responsible for different hardware parts, and is independent.
Apart from all helper functions, which are localed inside `utils`.
