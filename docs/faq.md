# FAQ

<details>
<summary>I'm seeing a lot of <code>info</code> level logs with the error <code>"404:Not Found - waiting for wormhole to pickup tx".</code></summary>

This is expected and OK ✅.

Jester is waiting for Wormhole to pick up the transaction which requires finality on Ethereum.
</details>

---

<details>
<summary>What happens if I'm a Validator and I don't run Jester?</summary>

You will see the following error each time you propose a new block:

```log
8:49AM ERR failed to get vote extension from Jester. Ensure Jester is configured and running! err="unavailable: dial tcp [::1]:9091: connect: connection refused" module=server
```

Currently, there are no other consequences.
</details>
