<script>
  import { Router, Route, Link } from "svelte-routing";
  import Login from "./routes/Login.svelte";
  import SignUp from "./routes/SignUp.svelte";
  import Home from "./routes/Home.svelte";
</script>

<div>
  <Router>
    <div>
      <Route path="/login" let:params component={Login} />
      <Route path="/signup" let:params component={SignUp} />
      <Route path="/" let:params component={Home} />
    </div>
  </Router>
  <main></main>
</div>
