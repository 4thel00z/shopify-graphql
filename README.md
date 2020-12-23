# shopify-graphql

## Motivation

I want to reverse engineer the shopify graphql API.

## What is the approach

Using this:

```
query IntrospectionQuery {
  __schema {
    queryType {
      name
    }
    mutationType {
      name
    }
    subscriptionType {
      name
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
    }
  }
}

fragment FullType on __Type {
  kind
  name
  description
  fields(includeDeprecated: true) {
    name
    description
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
}

fragment InputValue on __InputValue {
  name
  description
  type {
    ...TypeRef
  }
  defaultValue
}

fragment TypeRef on __Type {
  kind
  name
  ofType {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
              }
            }
          }
        }
      }
    }
  }
}
```
in the [graphql explorer](https://shopify.dev/tools/graphiql-admin-api) you can extract the `schema.json`.
Find mine [here](./schema.json). (I have remove the unnecessary parts that are returned by the explorer).

Using [graphql-introspection-json](https://github.com/potatosalad/graphql-introspection-json-to-sdl) you can easily generate the sdl from it using:

```
graphql-introspection-json-to-sdl schema.json > schema.graphql
```

Last step is to generate the stubs via [gqlgen](https://github.com/99designs/gqlgen).

I spent the first half hours chasing a phantom bug which I assumed in gqlgen. However you have to remove the first alleged type inside the `schema.graphql`.
It is sth. like `QueryRoot` or sth...
If you don't throw it you will either:

* build the glqgen stuff yourself and debug it
* cry

Please don't cry. 

<p align="center">
<img src="./cry.jpg" alt="crying cat">
</p>

To spare you the cry etc. and because I am a nice person I will add the generated stuff here for reference.

## License

This project is licensed under the GPL-3 license.
