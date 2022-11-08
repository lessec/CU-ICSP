<script>
    import { Auth } from 'aws-amplify';
    import { UserStore } from '../Stores/user-store.js';
    import { navigate } from "svelte-routing";

    async function signOut() {
      try {
        await Auth.signOut().then(() => {
                  $UserStore = null;
                  navigate("/logout", {replace: true});
              });
      } catch (error) {
        console.log('error signing out: ', error);
      }
    }
</script>

<button on:click={signOut}>Sign Out</button>