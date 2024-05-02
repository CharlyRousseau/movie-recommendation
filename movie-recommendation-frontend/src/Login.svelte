<script lang="ts">
  let email = '';
  let password = '';

  function handleLogin() {
    console.log({email, password});
</script>

<form on:submit|preventDefault={handleLogin} class="flex flex-col p-4">
  <div class="mb-4">
    <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
    <input type="email" id="email" bind:value={email} class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" required>
  </div>

  <div class="mb-6">
    <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
    <input type="password" id="password" bind:value={password} class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" required>
  </div>

  <button type="submit" class="self-center px-4 py-2 bg-blue-500 hover:bg-blue-700 text-white font-bold rounded-md">Login</button>
</form>
