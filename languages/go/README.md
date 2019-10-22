# {{name}}

The purpose of this microservice is to create, revoke, edit and validate API keys.

## Key

Currently a key has the following format:

`TYPE_HASH`

where:

1. TYPE - The key's type. It can be PROD or DEV. Serves no purpose other than to indicate to the user what it was generated for.
2. HASH - An unsalted SHA256 hash of a UUIDv4. Since UUIDv4 has an entropy of 128 bits (safe to generate 2⁶¹ keys), salting is unnecessary and ensures we can safely hash each key and safely perform database lookups, as a key will always hash to the same string.

## Database

We store the following information in the database:

1. `id` - A unique identifier, used for revoking and editing
2. `key` - The key itself as described in the previous paragraph
3. `scopes` - A list of scopes the key has access to
4. `account` - The account this key belongs to
5. `title` - A user-defined title, used for identification purposes
6. `createdAt` - When the key was created
7. `updatedAt` - When the key was last updated. Defaults to the creation date
8. `deletedAt` - When the key was deleted. Defaults to null

## Considerations

1. Keys *cannot* be retrieved after being generated as this would impose serious security risks. It would be no different from e-mailing the user their password when they forget it. Keys have the same purpose as passwords and should be treated as such.
2. The timestamp fields try to be as accurate as possible, but this information is for debugging purposes only. This microsservice does not expose this information, as it is the history service's responsibility.
