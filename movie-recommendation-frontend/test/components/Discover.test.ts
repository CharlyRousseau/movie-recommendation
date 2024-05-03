import { test, suite, expect } from 'vitest'
import Discover from '../../src/routes/Discover.svelte'
import { tick } from 'svelte'
import sinon from 'sinon'

suite('Discover Component', () => {
    test('should render the component', async () => {
        const host = document.createElement('div')

        document.body.appendChild(host)

        const instance = new Discover({ target: host })

        expect(instance).toBeTruthy()

        await tick()

        expect(host.innerHTML).toContain('Filter')
    })
})