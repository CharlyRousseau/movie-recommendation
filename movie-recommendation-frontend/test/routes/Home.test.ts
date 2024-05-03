import { test, suite, expect } from 'vitest'
import Home from '../../src/routes/Home.svelte'
import { tick } from 'svelte'
import sinon from 'sinon'

suite('Home Component', () => {
    test('should render the component', async () => {
        const host = document.createElement('div')

        document.body.appendChild(host)

        const instance = new Home({ target: host })

        expect(instance).toBeTruthy()

        await tick()

        expect(host.innerHTML).toContain('Trending')
    })

    test('should fetch movies', async () => {
        const fetchStub = sinon.stub(global, 'fetch')
        fetchStub.resolves({
            json: () => Promise.resolve({ results: [], total_pages: 0 }),
        })

        const host = document.createElement('div')

        document.body.appendChild(host)

        const instance = new Home({ target: host })

        await tick()

        sinon.assert.calledOnce(fetchStub)
        fetchStub.restore()
    })
})