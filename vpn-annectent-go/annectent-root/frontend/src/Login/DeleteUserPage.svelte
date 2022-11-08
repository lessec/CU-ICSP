<script>
    import Button, { Label } from '@smui/button';
    import Logo from '../Assets/Logo.svelte';

    import { navigate } from "svelte-routing";
    import { Auth } from 'aws-amplify';

    async function deleteUser() {
        try {
            const result = await Auth.deleteUser();
            console.log(result);
            navigate("/", {replace: true});

        } catch (error) {
            console.log('Error deleting user', error);
        }
    }
</script>

<main>
    <div class="inner">
        <Logo />
        <div class="mdc-typography--headline2 mt">Your not allowed!</div>
        <div class="mdc-typography--subtitle1">You have to delete your accout.</div>
        <div class="mdc-typography--subtitle1">Or you have to contact Admin.</div>
        <Button color="secondary" on:click={deleteUser} variant="unelevated" style="min-width: 300px; margin-top: 10px">
            <Label>Delete Me</Label>
        </Button>
    </div>
</main>

<style>
    main {
        display: flex;
        justify-content: center;
    }
    div.inner {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .mt {
        margin-top: 40px;
    }
</style>