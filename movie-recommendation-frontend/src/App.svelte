<script>
  import { Router, Route, Link } from "svelte-routing";
  import Login from "./routes/Login.svelte";
  import SignUp from "./routes/SignUp.svelte";
  import Home from "./routes/Home.svelte";
  import Discover from "./routes/Discover.svelte";
  import Favorites from "./routes/Favorites.svelte";
</script>

<div>
  <Router>
    <div>
      <Route path="/login" let:params component={Login} />
      <Route path="/signup" let:params component={SignUp} />
      <Route path="/discover" let:params component={Discover} />
      <Route path="/favorites" let:params component={Favorites} />
      <Route path="/" let:params component={Home} />
    </div>
  </Router>
  <main></main>
</div>
