<template>
  <div
    id="radix-:r98:"
    role="alertdialog"
    aria-describedby="radix-:r9a:"
    aria-labelledby="radix-:r99:"
    data-state="open"
    class="fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] sm:rounded-lg"
    tabindex="-1"
    style="pointer-events: auto"
  >
    <div class="flex flex-col space-y-2 text-center sm:text-left">
      <h2 id="radix-:r99:" class="text-lg font-semibold">
        Are you absolutely sure?
      </h2>
      <p id="radix-:r9a:" class="text-sm text-muted-foreground">
        This action cannot be undone. This will permanently delete your account
        and remove your data from our servers.
      </p>
    </div>
    <div class="flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2">
      <button
        type="button"
        class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 mt-2 sm:mt-0"
      >
        Cancel</button
      ><button
        type="button"
        class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2"
      >
        Continue
      </button>
    </div>
  </div>
</template>
