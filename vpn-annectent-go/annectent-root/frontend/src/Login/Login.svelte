<script>
    import Textfield from '@smui/textfield';
    import HelperText from '@smui/textfield/helper-text';
    import Button, { Label } from '@smui/button';

    import { Auth } from 'aws-amplify';
    import { UserStore } from '../Stores/user-store.js';
    import { navigate } from "svelte-routing";
    import { CognitoHostedUIIdentityProvider } from '@aws-amplify/auth';
    import Logo from '../Assets/Logo.svelte';

    let username = '';
    let password = '';

    async function signIn() {
      try {
        $UserStore = await Auth.signIn(username, password);
        console.log($UserStore);
        if ($UserStore.challengeName === 'SOFTWARE_TOKEN_MFA') {
          console.log('SOFTWARE_TOKEN_MFA');
          let code = prompt("Please enter MFA code");
          const loggedUser = await Auth.confirmSignIn(
            $UserStore,   // Return object from Auth.signIn()
                code,   // Confirmation code  ß
                "SOFTWARE_TOKEN_MFA" // The challenge name
            ).then(data => {
                console.log(data);
                navigate("/", { replace: true });
            });
        } else {
          console.log('signIn success');
          navigate("/", { replace: true });
        }
      } catch (error) {
        console.log('error signing in', error);
      }
    }
  
    function goToRegister() {
      navigate("/register", { replace: true });
    }

    async function federatedSignInWithGoogle() {
      await Auth.federatedSignIn({provider: CognitoHostedUIIdentityProvider.Google }).then(
        postLogin
      );
    }
    async function federatedSignInWithAmazon() {
      await Auth.federatedSignIn({ provider: 'LoginWithAmazon' });
        postLogin
    }

    function postLogin() {
      try {
        const user = Auth.currentAuthenticatedUser();
        $UserStore = user;
        console.log(user);
        navigate("/", { replace: true });
      } catch (error) {
        console.log(error);
      }
    }
</script>

<main>
  <div>
    <div class="inner">
      <div class="mdc-typography--headline2 mt">Hi, there!</div>
      <div class="mdc-typography--subtitle1 mb">Please sign in to user our service.</div>
      <Logo class="logo" />
      <div class="margin-top-bottom">
        <Textfield variant="outlined" bind:value={username} label="User Name" style="min-width: 300px">
          <HelperText slot="helper">Enter your User Name</HelperText>
        </Textfield>
        <Textfield variant="outlined" bind:value={password} label="Password" type="password" style="min-width: 300px">
          <HelperText slot="helper">Enter your Password</HelperText>
        </Textfield>
        <Button color="secondary" on:click={signIn} variant="unelevated" style="min-width: 300px">
          <Label>Sign in with email</Label>
        </Button>
        <Button color="secondary" on:click={goToRegister} variant="outlined" style="min-width: 300px; margin-top: 10px">
          <Label>Sign Up</Label>
        </Button>
      </div>
      <hr size="2" width="30%"/>
      <div class="margin-top-bottom">
        <Button color="secondary" on:click={federatedSignInWithAmazon} variant="outlined" style="min-width: 300px">
          <Label>Sign in with Amazon</Label>
        </Button>
        <Button color="secondary" on:click={federatedSignInWithGoogle} variant="outlined" style="min-width: 300px; margin-top: 10px">
          <Label>Sign in with Google</Label>
        </Button>
      </div>
    </div>
  </div>
</main>

<style>
  main {
    display: flex;
    justify-content: center;
  }
  main > div {
    display: flex;
    justify-content: center;
  }
  div.inner {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  div.inner > * {
    display: flex;
    flex-direction: column;
  }
  .margin-top-bottom {
    margin-top: 30px;
    margin-bottom: 30px;
  }
  .mt {
        margin-top: 40px;
  }
  .mb {
        margin-bottom: 15px;
  }
</style>