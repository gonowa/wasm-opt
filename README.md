# wasm-opt
wasm optimizer as docker container (7.96MB image)

# motivation 
currently https://github.com/WebAssembly/binaryen doesn't provide windows binaries.
this is a `cross platform`(kinda) version of wasm-opt.

# requirements
 * docker
 * go

# install
```sh
go get -u github.com/gonowa/wasm-opt
```

# usage 

```
wasm-opt INFILE

Read, write, and optimize files

Options:
  --help,-h                                    Show this help message and exit
  --debug,-d                                   Print debug information to stderr
   -O                                          execute default optimization
                                               passes
   -O0                                         execute no optimization passes
   -O1                                         execute -O1 optimization passes
                                               (quick&useful opts, useful for
                                               iteration builds)
   -O2                                         execute -O2 optimization passes
                                               (most opts, generally gets most
                                               perf)
   -O3                                         execute -O3 optimization passes
                                               (spends potentially a lot of time
                                               optimizing)
   -O4                                         execute -O4 optimization passes
                                               (also flatten the IR, which can
                                               take a lot more time and memory,
                                               but is useful on more nested /
                                               complex / less-optimized input)
   -Os                                         execute default optimization
                                               passes, focusing on code size
   -Oz                                         execute default optimization
                                               passes, super-focusing on code
                                               size
  --optimize-level,-ol                         How much to focus on optimizing
                                               code
  --shrink-level,-s                            How much to focus on shrinking
                                               code size
  --no-validation,-n                           Disables validation, assumes
                                               inputs are correct
  --ignore-implicit-traps,-iit                 Optimize under the helpful
                                               assumption that no surprising
                                               traps occur (from load, div/mod,
                                               etc.)
  --coalesce-locals                            reduce # of locals by coalescing
  --coalesce-locals-learning                   reduce # of locals by coalescing
                                               and learning
  --code-folding                               fold code, merging duplicates
  --code-pushing                               push code forward, potentially
                                               making it not always execute
  --const-hoisting                             hoist repeated constants to a
                                               local
  --dce                                        removes unreachable code
  --duplicate-function-elimination             removes duplicate functions
  --extract-function                           leaves just one function (useful
                                               for debugging)
  --flatten                                    flattens out code, removing
                                               nesting
  --fpcast-emu                                 emulates function pointer casts,
                                               allowing incorrect indirect calls
                                               to (sometimes) work
  --func-metrics                               reports function metrics
  --i64-to-i32-lowering                        lower all uses of i64s to use
                                               i32s instead
  --inlining                                   inline functions (you probably
                                               want inlining-optimizing)
  --inlining-optimizing                        inline functions and optimizes
                                               where we inlined
  --instrument-locals                          instrument the build with code to
                                               intercept all loads and stores
  --instrument-memory                          instrument the build with code to
                                               intercept all loads and stores
  --legalize-js-interface                      legalizes i64 types on the
                                               import/export boundary
  --local-cse                                  common subexpression elimination
                                               inside basic blocks
  --log-execution                              instrument the build with logging
                                               of where execution goes
  --memory-packing                             packs memory into separate
                                               segments, skipping zeros
  --merge-blocks                               merges blocks to their parents
  --merge-locals                               merges locals when beneficial
  --metrics                                    reports metrics
  --nm                                         name list
  --optimize-instructions                      optimizes instruction
                                               combinations
  --pick-load-signs                            pick load signs based on their
                                               uses
  --post-emscripten                            miscellaneous optimizations for
                                               Emscripten-generated code
  --precompute                                 computes compile-time evaluatable
                                               expressions
  --precompute-propagate                       computes compile-time evaluatable
                                               expressions and propagates them
                                               through locals
  --print                                      print in s-expression format
  --print-call-graph                           print call graph
  --print-full                                 print in full s-expression format
  --print-minified                             print in minified s-expression
                                               format
  --relooper-jump-threading                    thread relooper jumps (fastcomp
                                               output only)
  --remove-imports                             removes imports and replaces them
                                               with nops
  --remove-memory                              removes memory segments
  --remove-non-js-ops                          removes operations incompatible
                                               with js
  --remove-unused-brs                          removes breaks from locations
                                               that are not needed
  --remove-unused-module-elements              removes unused module elements
  --remove-unused-names                        removes names from locations that
                                               are never branched to
  --remove-unused-nonfunction-module-elements  removes unused module elements
                                               that are not functions
  --reorder-functions                          sorts functions by access
                                               frequency
  --reorder-locals                             sorts locals by access frequency
  --rereloop                                   re-optimize control flow using
                                               the relooper algorithm
  --rse                                        remove redundant set_locals
  --safe-heap                                  instrument loads and stores to
                                               check for invalid behavior
  --simplify-locals                            miscellaneous locals-related
                                               optimizations
  --simplify-locals-nonesting                  miscellaneous locals-related
                                               optimizations (no nesting at all;
                                               preserves flatness)
  --simplify-locals-nostructure                miscellaneous locals-related
                                               optimizations
  --simplify-locals-notee                      miscellaneous locals-related
                                               optimizations
  --simplify-locals-notee-nostructure          miscellaneous locals-related
                                               optimizations
  --spill-pointers                             spill pointers to the C stack
                                               (useful for Boehm-style GC)
  --ssa                                        ssa-ify variables so that they
                                               have a single assignment
  --trap-mode-clamp                            replace trapping operations with
                                               clamping semantics
  --trap-mode-js                               replace trapping operations with
                                               js semantics
  --untee                                      removes tee_locals, replacing
                                               them with sets and gets
  --vacuum                                     removes obviously unneeded code
  --output,-o                                  Output file (stdout if not
                                               specified)
  --emit-text,-S                               Emit text instead of binary for
                                               the output file
  --debuginfo,-g                               Emit names section and debug info
  --converge,-c                                Run passes to convergence,
                                               continuing while binary size
                                               decreases
  --fuzz-exec,-fe                              Execute functions before and
                                               after optimization, helping
                                               fuzzing find bugs
  --fuzz-binary,-fb                            Convert to binary and back after
                                               optimizations and before
                                               fuzz-exec, helping fuzzing find
                                               binary format bugs
  --extra-fuzz-command,-efc                    An extra command to run on the
                                               output before and after
                                               optimizing. The output is
                                               compared between the two, and an
                                               error occurs if they are not
                                               equal
  --translate-to-fuzz,-ttf                     Translate the input into a valid
                                               wasm module *somehow*, useful for
                                               fuzzing
  --fuzz-passes,-fp                            Pick a random set of passes to
                                               run, useful for fuzzing. this
                                               depends on translate-to-fuzz (it
                                               picks the passes from the input)
  --emit-js-wrapper,-ejw                       Emit a JavaScript wrapper file
                                               that can run the wasm with some
                                               test values, useful for fuzzing
  --emit-spec-wrapper,-esw                     Emit a wasm spec interpreter
                                               wrapper file that can run the
                                               wasm with some test values,
                                               useful for fuzzing
  --input-source-map,-ism                      Consume source map from the
                                               specified file
  --output-source-map,-osm                     Emit source map to the specified
                                               file
  --output-source-map-url,-osu                 Emit specified string as source
                                               map URL

```

