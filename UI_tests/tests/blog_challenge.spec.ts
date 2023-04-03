import { test, expect } from '@playwright/test';

test('has title', async ({ page }) => {
  await page.goto('http://localhost:3000/');

  // Expect a title "to contain" a substring.
  await expect(page).toHaveTitle(/Blog Challenge with Markdown/);
});

test('has expected blogpost', async ({ page }) => {
  await page.goto('http://localhost:3000/');

  // Click the get started link.
  await page.getByRole('link', { name: 'Learn How to Pre-render Pages Using Static Generation with Next.js' }).first().click();

  // Expects the URL to contain intro.
  await expect(page).toHaveURL(/.*hello-world/);
});
