# ddwrapper
A wrapper for the DataDog client

This package simply wraps the DataDog client created by the statsd package to provide some convenient usage.

- []String Tags are converted datadog proper tags key:value paring, e.g []string{"hello", "world"} -> []string{"hello:world}
    - a special note, if the supplied []string is not even, paring will not happen. All Tags in datadog require the slice to be even
    and just like the statsd package, the caller is responsible for ensuring tags are proper.
- FixRate methods where added to avoid having to supply a rate of 1 constantly, making the API nicer to work with.  