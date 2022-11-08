<script>
    import Textfield from '@smui/textfield';
    import HelperText from '@smui/textfield/helper-text';
    import Button, { Label } from '@smui/button';

    import { Auth } from 'aws-amplify';
    import {UserStore} from '../Stores/user-store.js';
    import { navigate } from "svelte-routing";
    import Logo from '../Assets/Logo.svelte';

    let username = '';
    let password = '';
    let email = '';

    async function signUp() {
		try {
			await Auth.signUp({
				username,
				password,
				attributes: {
                    email,
                    'custom:usertype': "Staff"
				}
			}).then(() => {
                
                $UserStore = {
                    name: username,
                    password,
                    email: email,
                    type: "Staff"
                };
                console.log($UserStore)
                navigate("/confirm", { replace: true });
            });
		} catch (error) {
			console.log('error signing up:', error);
		}
	}
</script>

<main>
    <div>
        <div class="inner">
          <div class="mdc-typography--headline2 mt">Wellcome!</div>
          <div class="mdc-typography--subtitle1 mb">Nice to meet you!</div>
          <Logo class="logo" />
          <div class="margin-top-bottom">
            <Textfield variant="outlined" bind:value={username} label="User Name" style="min-width: 300px">
              <HelperText slot="helper">Enter your User Name</HelperText>
            </Textfield>
            <Textfield variant="outlined" bind:value={password} label="Password" type="password" style="min-width: 300px">
              <HelperText slot="helper">Enter your Password</HelperText>
            </Textfield>
            <Textfield variant="outlined" bind:value={email} label="Email" type="email" style="min-width: 300px">
                <HelperText slot="helper">Enter your Email</HelperText>
            </Textfield>
            <Button color="secondary" on:click={signUp} variant="unelevated" style="min-width: 300px">
              <Label>Sign in with email</Label>
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
    margin-top: 50px;
    margin-bottom: 30px;
  }
  .mt {
        margin-top: 40px;
  }
  .mb {
        margin-bottom: 15px;
  }
</style>