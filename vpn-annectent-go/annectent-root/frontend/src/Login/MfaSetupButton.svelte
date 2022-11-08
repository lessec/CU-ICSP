<script>
    import { navigate } from "svelte-routing";
</script>
<button on:click={()=>{navigate("/mfa", {replace: true})}}>MFA</button>