import { test, suite, expect } from 'vitest'
import UserAuthFormSignup from '../../src/components/user-auth-form-signup.svelte'
import { tick } from 'svelte'

suite('UserAuthFormSignup Component', () => {
    test('should render the component', async () => {
        const host = document.createElement('div')

        document.body.appendChild(host)

        const instance = new UserAuthFormSignup({ target: host })

        expect(instance).toBeTruthy()

        await tick()

        expect(host.innerHTML).toContain('Sign Up')
    })

    test('should set isLoading to true on form submission', async () => {
        const host = document.createElement('div')

        document.body.appendChild(host)

        const instance = new UserAuthFormSignup({ target: host })

        await tick()

        const form = host.querySelector('form')
        if (!form) throw new Error('Form not found')
        form.dispatchEvent(new Event('submit'))

        await tick()

        expect(instance.$$.ctx[4]).toBe(true)
    })
})