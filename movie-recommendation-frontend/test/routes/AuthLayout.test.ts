import { test, suite, expect, describe } from 'vitest'
import AuthLayout from '../../src/routes/AuthLayout.svelte'

import { render, screen } from '@testing-library/svelte';

describe("VerticalTabs Component", () => {

  test("should render the component", () => {

    render(AuthLayout);

    const firstTabNode = screen.getByText(/Login/i)

    expect(firstTabNode).toBeTruthy()
  });

})