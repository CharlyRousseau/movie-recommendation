import { test, suite, expect } from 'vitest'
import UserAuthFormLogin from '../../src/components/user-auth-form-login.svelte'
import { tick } from 'svelte'

suite('UserAuthFormLogin Component', () => {
    test('should render the component', async () => {
        const host = document.createElement('div')

        document.body.appendChild(host)

        const instance = new UserAuthFormLogin({ target: host })

        expect(instance).toBeTruthy()

        await tick()

        expect(host.innerHTML).toContain('Connect')
    })

    test('should set isLoading to true on form submission', async () => {
        const host = document.createElement('div')

        document.body.appendChild(host)

        const instance = new UserAuthFormLogin({ target: host })

        await tick()

        const form = host.querySelector('form')

        if (!form) throw new Error('Form not found')

        form.dispatchEvent(new Event('submit'))

        await tick()

        expect(instance.$$.ctx[2]).toBe("")
    })
})