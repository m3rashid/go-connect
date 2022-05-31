import React from "react";
import {
  useFilters,
  useGlobalFilter,
  usePagination,
  useSortBy,
  useTable,
} from "react-table";
import {
  MdChevronLeft,
  MdChevronRight,
  MdDoubleArrow,
  MdOutlineArrowDropDown,
  MdOutlineArrowDropUp,
  MdSort,
} from "react-icons/md";
import Select from "react-select";

import { Button, PageButton, GlobalFilter } from "./tableHelpers";

export const Table = ({ columns, data, title }) => {
  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    prepareRow,
    page,
    canPreviousPage,
    canNextPage,
    pageOptions,
    pageCount,
    gotoPage,
    nextPage,
    previousPage,
    setPageSize,
    state,
    preGlobalFilteredRows,
    setGlobalFilter,
  } = useTable(
    { columns, data },
    useFilters,
    useGlobalFilter,
    useSortBy,
    usePagination
  );

  React.useEffect(() => setPageSize(5), [setPageSize]);
  const { pageSize, pageIndex } = state;
  const handlePageChange = (selected) => setPageSize(selected.value);

  return (
    <div className="flex flex-col mt-10 justify-between bg-gray-100 dark:bg-gray-700 rounded-md w-full max-w-[1400px] min-h-[500px]">
      <div className="flex flex-col flex-grow">
        <div className="p-5 pb-2 flex justify-between">
          <h1 className="font-bold text-2xl dark:text-gray-50">{title}</h1>
          <GlobalFilter
            preGlobalFilteredRows={preGlobalFilteredRows}
            globalFilter={state.globalFilter}
            setGlobalFilter={setGlobalFilter}
          />
          {headerGroups.map((headerGroup) =>
            headerGroup.headers.map((column) =>
              column.Filter ? (
                <div className="mt-2 sm:mt-0" key={column.id}>
                  {column.render("Filter")}
                </div>
              ) : null
            )
          )}
        </div>
        <div className="mt-4 flex flex-col w-full">
          <div className="-my-2 overflow-x-auto">
            <div className="py-2 align-middle inline-block min-w-full">
              <div className="overflow-hidden">
                <table
                  {...getTableProps()}
                  className="min-w-full divide-y divide-gray-500 dark:divide-gray-500"
                >
                  <thead className="bg-gray-50 dark:bg-gray-900">
                    {headerGroups.map((headerGroup) => (
                      <tr {...headerGroup.getHeaderGroupProps()}>
                        {headerGroup.headers.map((column) => (
                          <th
                            scope="col"
                            className="group px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-100 uppercase tracking-wider"
                            {...column.getHeaderProps(
                              column.getSortByToggleProps()
                            )}
                          >
                            <div className="flex items-center justify-between">
                              {column.render("Header")}
                              <span>
                                {column.isSorted ? (
                                  column.isSortedDesc ? (
                                    <MdOutlineArrowDropDown className="w-4 h-4 text-gray-400" />
                                  ) : (
                                    <MdOutlineArrowDropUp className="w-4 h-4 text-gray-400" />
                                  )
                                ) : (
                                  <MdSort className="w-4 h-4 text-gray-400 opacity-0 group-hover:opacity-100" />
                                )}
                              </span>
                            </div>
                          </th>
                        ))}
                      </tr>
                    ))}
                  </thead>
                  <tbody
                    {...getTableBodyProps()}
                    className="bg-gray-50 dark:bg-gray-700 divide-y divide-gray-500 dark:divide-gray-100"
                  >
                    {page.map((row, i) => {
                      prepareRow(row);
                      return (
                        <tr {...row.getRowProps()}>
                          {row.cells.map((cell) => {
                            return (
                              <td
                                {...cell.getCellProps()}
                                className="px-6 py-4 whitespace-nowrap"
                                role="cell"
                              >
                                <div className="dark:text-gray-50">
                                  {cell.render("Cell")}
                                </div>
                              </td>
                            );
                          })}
                        </tr>
                      );
                    })}
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="px-5 py-3 flex items-center justify-between">
        <div className="flex-1 flex justify-between sm:hidden">
          <Button
            className="dark:text-gray-50 cursor-pointer"
            onClick={() => previousPage()}
            disabled={!canPreviousPage}
          >
            Previous
          </Button>
          <Button
            className="dark:text-gray-50 cursor-pointer"
            onClick={() => nextPage()}
            disabled={!canNextPage}
          >
            Next
          </Button>
        </div>
        <div className="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div className="flex gap-x-2 items-baseline">
            <span className="text-sm dark:text-gray-50 whitespace-nowrap">
              Page {pageIndex + 1} of {pageOptions.length}
            </span>
            <Select
              name=""
              options={[5, 10, 15, 20].map((n) => ({ label: n, value: n }))}
              placeholder={`${pageSize} items`}
              onChange={handlePageChange}
              value={pageSize}
              single={true}
              styles={{
                control: (base) => ({
                  ...base,
                  border: "none",
                  minWidth: "150px",
                }),
              }}
              classNamePrefix="bg-gray-200 dark:bg-gray-700 text-black dark:text-white outline-none"
            />
          </div>
          <div>
            <nav
              className="relative z-0 inline-flex rounded-md shadow-sm -space-x-px"
              aria-label="Pagination"
            >
              <PageButton
                className="rounded-l-md"
                onClick={() => gotoPage(0)}
                disabled={!canPreviousPage}
              >
                <span className="sr-only">First</span>
                <MdDoubleArrow
                  className="h-5 w-5 text-gray-400 rotate-180"
                  aria-hidden="true"
                />
              </PageButton>
              <PageButton
                onClick={() => previousPage()}
                disabled={!canPreviousPage}
              >
                <span className="sr-only">Previous</span>
                <MdChevronLeft
                  className="h-5 w-5 text-gray-400"
                  aria-hidden="true"
                />
              </PageButton>
              <PageButton onClick={() => nextPage()} disabled={!canNextPage}>
                <span className="sr-only">Next</span>
                <MdChevronRight
                  className="h-5 w-5 text-gray-400"
                  aria-hidden="true"
                />
              </PageButton>
              <PageButton
                className="rounded-r-md"
                onClick={() => gotoPage(pageCount - 1)}
                disabled={!canNextPage}
              >
                <span className="sr-only">Last</span>
                <MdDoubleArrow
                  className="h-5 w-5 text-gray-400"
                  aria-hidden="true"
                />
              </PageButton>
            </nav>
          </div>
        </div>
      </div>
    </div>
  );
};
